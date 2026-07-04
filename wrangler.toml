#:schema node_modules/wrangler/config-schema.json
# 👑 มหาอาณาจักร TNH AI V84.5 TRINITY EMPIRE ENGINE 
# [ARCH: ZERO-GARBAGEQUEUES] [SECURITY: FIXED BINDINGS] [PORT: 2026]

name = "tnh-ai-v83-trinity-empire" # ล็อกชื่อแท้ตรงตามระบบ Cloudflare CI เรียบร้อย ไร้ Warning
main = "src/index.ts"
compatibility_date = "2026-07-04"
compatibility_flags = [ "nodejs_compat" ]

# --- [WASM RULES CONFIGURATION] ---
[[rules]]
type = "CompiledWasm"
globs = ["**/*.wasm"]
fallthrough = true

# ==========================================================
# [RESOURCE BINDINGS - 11 GUARDIANS CORE RESOURCE]
# ==========================================================

# 1. ฐานข้อมูล D1 สัจจะข้อมูล (ตารางสิทธิ์ 11 ขุนพลเสือขาว)
[[d1_databases]]
binding = "DB" 
database_name = "thitnueahub-core-db"
database_id = "6a8b4373-bf40-4b63-bb02-f612ecbe63b7"

# 2. คลังจัดเก็บไฟล์ R2 Buckets สำหรับสตรีมมิ่งภาพและโมเดลสายตา (.safetensors)
[[r2_buckets]]
binding = "TNH_R2"
bucket_name = "thitnueahub-assets"

# 3. คลัง KV จัดเก็บไทม์สแตมป์ความจำถาวร 20 ปี ป้องกันบอทหลับ
[[kv_namespaces]]
binding = "TNH_KV"
id = "2fa0a4773efa4a18b1534274d238dd76"

# 4. โมเดล AI Edge ขอบสนาม (Edge Intelligence LlaVA)
[ai]
binding = "AI"

# ==========================================================
# [CLOUDFLARE QUEUES OMNI-CHANNEL PIPELINE]
# ยึดท่อระบายข้อมูลฝั่งคลาวด์แฟลร์ tnh-files.9thera.workers.dev ปิดตายกูเกิ้ล
# ==========================================================

# ท่อส่งออกคำสั่ง (Producers) - สำหรับ L1 Commander สั่งการ หรือบอทลูกน้องยิงสลับหน้าที่
[[queues.producers]]
queue = "tnh-orchestrator-queue"
binding = "ORCHESTRATOR_QUEUE"

# ท่อประมวลผลขาเข้า (Consumers) - มีระบบป้องกันข้อมูลหลุด ประมวลผลล้มเหลวจะเข้าตู้ DLQ อัตโนมัติ
[[queues.consumers]]
queue = "tnh-orchestrator-queue"
max_batch_size = 10         # ดึงข้อมูลรันทีละ 10 สัญญะ ป้องการหน่วยความจำล้น (Zero-Garbage)
max_batch_timeout = 5       # รอรวมคิวสูงสุด 5 วินาที สับไวระดับไมโครวินาที
max_retries = 3             # รันซ้ำสูงสุด 3 รอบ ถ้าพังเกินจะดีดเข้า DLQ
dead_letter_queue = "tnh-garbage-dlq" # ตู้เซฟกักขยะข้อมูล ให้ขุนพล L5 ไอ้จ๊อด ไปล้างบาง

# เปิดระบบสังเกตการณ์พฤติกรรม (Observability) ตามพิมพ์เขียวแท้ (🌱 j + ทอม 👾)
[observability.traces]
enabled = true
destinations = [ "tnh-tracing-hub" ]
head_sampling_rate = 0.05
persist = false

[observability.logs]
enabled = true
destinations = [ "tnh-log-sink" ]
head_sampling_rate = 0.60
persist = false
