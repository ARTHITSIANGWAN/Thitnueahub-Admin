import time
import logging
import subprocess
import os
import google.generativeai as genai
from dotenv import load_dotenv

# โหลดค่าคอนฟิกจากไฟล์ .env (สำหรับใส่ API Key)
load_dotenv()

# ตั้งค่าระบบแจ้งเตือน (Logger)
logging.basicConfig(level=logging.INFO, format='%(asctime)s - [KAEWTA] - %(message)s')
logger = logging.getLogger(__name__)

class KaewtaSupervisorLinux:
    def __init__(self):
        logger.info(" 😉 แก้วตา (Linux/Termux Native) เริ่มเดินตรวจระบบ...")
        # ใส่ชื่อไฟล์สคริปต์ที่บอสต้องการให้แก้วตาเฝ้าดู
        self.agents = ["james_chatbot.py", "nong_ing_backend.js"]
        
        # ติดตั้งสมองกล (Gemini AI) ให้แก้วตา
        api_key = os.getenv("GEMINI_API_KEY")
        if api_key:
            genai.configure(api_key=api_key)
            self.ai_model = genai.GenerativeModel('gemini-1.5-flash')
            logger.info("🐍 แก้วตาเชื่อมต่อสมองกล AI สำเร็จแล้ว! พร้อมวิเคราะห์ปัญหา")
        else:
            self.ai_model = None
            logger.warning(" ไม่พบ GEMINI_API_KEY ในระบบ แก้วตาจะทำงานแบบเช็กสถานะทั่วไป (ไม่มี AI)")

    def check_process(self, script_name):
        """ตรวจสอบว่าโปรเซสกำลังทำงานอยู่บน Linux หรือไม่"""
        try:
            # ใช้คำสั่ง Linux พื้นฐานค้นหาชื่อไฟล์ที่กำลังรัน
            cmd = f"ps -ef | grep '{script_name}' | grep -v grep"
            output = subprocess.check_output(cmd, shell=True, text=True)
            return True if output.strip() else False
        except subprocess.CalledProcessError:
            # ถ้าหาไม่เจอ คำสั่งจะพ่น Error ทำให้รู้ว่าโปรแกรมตาย
            return False

    def heal_process(self, agent_name):
        """ชุบชีวิตระบบ และวิเคราะห์ปัญหาด้วย AI"""
        logger.warning(f" 🫧 ตรวจพบ {agent_name} หยุดทำงาน! แก้วตากำลังเตรียมชุบชีวิต...")
        
        # ให้ AI ช่วยวิเคราะห์สาเหตุเบื้องต้น
        if self.ai_model:
            prompt = f"โปรแกรมชื่อ {agent_name} บน Linux/Termux พังและหยุดทำงาน ในฐานะผู้คุมระบบ (Supervisor) ช่วยแนะนำสั้นๆ 1-2 ประโยค ว่าควรเช็กที่จุดไหนเป็นอันดับแรก"
            try:
                response = self.ai_model.generate_content(prompt)
                logger.info(f" 🐥 วิเคราะห์จากแก้วตา: {response.text.strip()}")
            except Exception as e:
                logger.error(f"เชื่อมต่อ AI มีปัญหา: {e}")

        # เริ่มต้นรันโปรแกรมใหม่ทิ้งไว้ใน Background
        try:
            # ให้บันทึก Log ของแต่ละ Agent แยกไว้ด้วย จะได้ตามสืบได้
            log_file = f"{agent_name.split('.')[0]}_error.log"
            
            if agent_name.endswith('.py'):
                os.system(f"nohup python3 {agent_name} > {log_file} 2>&1 &")
            elif agent_name.endswith('.js'):
                os.system(f"nohup node {agent_name} > {log_file} 2>&1 &")
                
            logger.info(f" ♻️ ชุบชีวิต {agent_name} สำเร็จ! (ระบบจะบันทึก Log ไว้ที่ {log_file})")
        except Exception as e:
            logger.error(f"ชุบชีวิต {agent_name} ไม่สำเร็จ: {e}")

    def run(self):
        logger.info(" -🧐-  ระบบ ThitNuea Hub (Linux Edge Edition) พร้อมลุย ---")
        while True:
            for agent in self.agents:
                if self.check_process(agent):
                    # ถ้าระบบปกติ จะไม่แสดงข้อความเพื่อไม่ให้รกหน้าจอ (Zero-Garbage Output)
                    pass
                else:
                    self.heal_process(agent)
            
            # แก้วตาเดินตรวจตราทุกๆ 30 วินาที
            time.sleep(30)

if __name__ == "__main__":
    supervisor = KaewtaSupervisorLinux()
    supervisor.run()

