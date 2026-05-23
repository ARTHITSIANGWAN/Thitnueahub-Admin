import { WorkflowEntrypoint, WorkflowStep } from "cloudflare:workers";
import type { WorkflowEvent } from "cloudflare:workers";

// 1. ENVIRONMENT INTERFACE DEFINITION (ZERO TRUST BINDINGS)
// [แก้ไข]: ปรับชื่อจาก BUCKET เป็น TNH_R2 ให้ตรงกับไฟล์ wrangler.toml เป๊ะๆ ป้องกันระบบหาบัคเก็ตไม่เจอแล้วเอ๋อ
export interface Env {
  TNH_KV: KVNamespace;
  DB: D1Database;
  TNH_R2: R2Bucket; 
  AI: Ai;
  MY_WORKFLOW: WorkflowNamespace;
}

type ImageParams = { imageKey: string };

// 2. DURABLE EXECUTION WORKFLOW ENGINE (L3/L5 AGENT CORE)
export class ImageProcessingWorkflow extends WorkflowEntrypoint<Env, ImageParams> {
  async run(event: WorkflowEvent<ImageParams>, step: WorkflowStep) {
    
    // [ปฏิวัติโครงสร้าง Step A & B]: ยุบรวมการดึงภาพและการยิง AI ให้อยู่ใน Step เดียวกัน
    // เพื่อป้องกันไม่ให้ระบบส่งก้อน Array ของตัวเลขฐานสิบหลายล้านตัวข้ามสเต็ป 
    // ซึ่งช่วยลดการสร้างขยะ JSON ขนาดใหญ่ในฐานข้อมูลคลาวด์ (State Database Data Waste)
    const description = await step.do("analyze image via AI", async () => {
      const object = await this.env.TNH_R2.get(event.payload.imageKey);
      if (!object) {
        throw new Error(`[FAIL] Object metadata missing: ${event.payload.imageKey}`);
      }
      
      const buffer = await object.arrayBuffer();
      const uint8Data = new Uint8Array(buffer);

      // รันโมเดลสายตาประมวลผลในแรมทันที แล้วสลายตัวแปรทิ้ง ไม่ทิ้งซากขยะไว้ในระบบ
      return await this.env.AI.run("@cf/llava-hf/llava-1.5-7b-hf", {
        image: Array.from(uint8Data),
        prompt: "Describe this image in one sentence",
        max_tokens: 50,
      });
    });

    // STEP C: Human-in-the-loop lock (ด่านกักตรวจ รอคำสั่งอนุมัติจากบอสภายใน 24 ชั่วโมง)
    await step.waitForEvent("await approval", {
      event: "approved",
      timeout: "24 hours",
    });

    // [ปฏิวัติโครงสร้าง Step D]: ใช้ระบบ Streaming ข้อมูลข้ามโฟลเดอร์ใน R2 โดยตรง 
    // ไม่โหลดไฟล์ทั้งก้อนลงหน่วยความจำซ้ำซ้อน ช่วยให้ประหยัด CPU และแรมของบอร์ดคลาวด์ขั้นสูงสุด
    await step.do("publish to public layer", async () => {
      const object = await this.env.TNH_R2.get(event.payload.imageKey);
      if (!object) {
        throw new Error(`[FAIL] Object disappeared before publish: ${event.payload.imageKey}`);
      }
      
      // ส่งต่อผ่าน object.body (ReadableStream) ทันที คลีนระดับโมเลกุล
      await this.env.TNH_R2.put(`public/${event.payload.imageKey}`, object.body);
    });

    return { status: "SUCCESS", description: description };
  }
}

// 3. MAIN GATEWAY ROUTER (FETCH HANDLER)
export default {
  async fetch(request: Request, env: Env, ctx: any): Promise<Response> {
    const url = new URL(request.url);

    // [ROUTE A]: TRIGGER NEW IMAGE WORKFLOW VIA API
    if (url.pathname === "/workflow/start") {
      const imageKey = url.searchParams.get("key") || "matrix-vision.jpg";
      const instance = await env.MY_WORKFLOW.create({
        params: { imageKey: imageKey }
      });
      return Response.json({ status: "WORKFLOW_LAUNCHED", instanceId: instance.id });
    }

    // [ROUTE B]: UNSTUCK GATEWAY - ส่งสัญญาณปลดล็อกด่านกักตรวจให้ Workflow เดินหน้าต่อ
    if (url.pathname === "/workflow/approve") {
      const instanceId = url.searchParams.get("instanceId");
      if (!instanceId) return new Response("[FAIL] Missing instanceId", { status: 400 });
      
      const instance = await env.MY_WORKFLOW.get(instanceId);
      await instance.sendEvent({ type: "approved", payload: {} });
      return Response.json({ status: "EVENT_DISPATCHED", target: instanceId });
    }

    // [DEFAULT ROUTE]: คอนโซลเช็คสถานะหลักของจักรวรรดิ
    await env.TNH_KV.put('LAST_BOOT', new Date().toISOString());

    // ตรวจเช็คความเสถียรของดีบี D1
    const { results } = await env.DB.prepare("SELECT name FROM sqlite_master LIMIT 1").all();

    return new Response(
      JSON.stringify({
        status: "TNH-AI-V83-ONLINE",
        message: "อาณาจักร 9THERA ปลดล็อกป้ายชื่อตรงกันเรียบร้อยก้า! 🤫",
        kv_check: await env.TNH_KV.get('LAST_BOOT'),
        db_ready: !!results,
        path_requested: url.pathname
      }),
      { 
        headers: { 
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
          "Content-Language": "en" 
        } 
      }
    );
  }
};
              
