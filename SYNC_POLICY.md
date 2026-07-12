# 🏰 TNH AI MASTER: SYNC POLICY & PROTOCOL
## 🌱 Conducting the Future © 2026 9THERA EMPIRE 🫟

### 1. Vision & Governance
เอกสารฉบับนี้คือ "กฎเหล็ก" สำหรับการประสานงานระหว่าง **Hub (TNH_AI_MASTER)** และ **Spokes (12 Projects)**. โปรเจกต์หลักทำหน้าที่เป็นศูนย์กลางเก็บ DNA และมาตรฐาน ระบบ Sync ต้องทำงานเพื่อรักษาความสม่ำเสมอของโครงสร้าง โดยห้ามส่งผลกระทบต่อความเสถียรของงานเฉพาะทางในแต่ละสาขา

### 2. 4 Quadrants Foundation (DNA Alignment)
ทุก Spoke ต้องคงโครงสร้าง 4 ส่วนนี้ไว้เสมอ:
- **Identity (Q1):** โครงสร้างตัวตนที่สืบทอดจาก Master
- **Operation (Q2):** Workflow A, B, C (มาตรฐานการรันงาน)
- **Governance (Q3):** FiveTwo.ml (ตัวกรองความสะอาด)
- **Memory (Q4):** Chronicle.ml (อาลักษณ์บันทึกผลงาน)

### 3. Sync Rules & Protocols
- **Source of Truth:** Master คือแหล่งข้อมูลที่ถูกต้องที่สุด หากเกิดความขัดแย้ง (Conflict) ให้ยึดจาก Master เป็นหลัก
- **Sync Direction:** Hub -> Spokes เท่านั้น (ห้าม Spoke อัปเดต Hub โดยตรง เพื่อป้องกัน Poisoning)
- **Versioning:** การ Sync ต้องอ้างอิงผ่าน `Version Tag` (เช่น `v1.0.0-Stable`) เสมอ ห้ามดึงข้อมูลแบบ `main` หรือ `master` ลอยๆ
- **Forward Progress:** ทุกการ Sync ต้องไม่ทำให้งานใน Spoke ต้อง Restart (ใช้ Micro-Correction ผ่าน Checkpoints)

### 4. White/Black List (Hard Constraints)
| Action | Included (Sync ทุกครั้ง) | Excluded (ห้ามยุ่งเด็ดขาด) |
| :--- | :--- | :--- |
| **Files** | `Context.ml`, `FiveTwo.ml`, `Standard_Workflow_A.md` | `.env`, `.secrets`, `Local_DB`, `Local_Config` |
| **Logic** | Core Engine updates, Security patches | Local specialized features |

### 5. Verification Protocol
ทุกครั้งที่ Sync สำเร็จ `Chronicle.ml` ใน Spoke ต้องส่งสัญญาณ "Verify" กลับมาที่ Master เพื่อยืนยันว่าการอัปเดตสมบูรณ์และระบบยังทำงานปกติ (Self-Healing Check)

---
*ความผิดพลาดใดๆ ที่เกิดจากความไม่ปฏิบัติตาม Protocol นี้ถือเป็นความบกพร่องของระบบ Governance ในโปรเจกต์นั้นๆ*
