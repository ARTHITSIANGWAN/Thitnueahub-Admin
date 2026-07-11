package registry

type AgentProfile struct {
	Level    string
	Role     string
	Name     string
	Function string
	Base     string // ตัวตนพื้นฐาน (Base Role)
}

var Registry = map[string]AgentProfile{
	"L1":  {"L1", "COMMANDER", "ทิศเหนือ", "ผู้บัญชาการสูงสุด (The Conductor)", "Jarvis System"},
	"L2":  {"L2", "SCRIBE", "แก้วตา", "ผู้ดูแลสัจจะและการสื่อสารกับลูกค้า", "Jarvis System"},
	"L3":  {"L3", "ARTIST", "น้ำอิง", "ผู้สร้างสรรค์และศัลยกรรมตรรกะ (PNG encoding)", "Jarvis System"},
	"L4":  {"L4", "ANALYST", "พลายทอง", "นักวิเคราะห์ Edge Intelligence", "Jarvis System"},
	"L5":  {"L5", "AUDITOR", "ไอ้จ๊อด", "ผู้คุมกฎ Zero-Garbage", "Jarvis System"},
	"L6":  {"L6", "MERCHANT", "Merchant", "ฝ่ายพัฒนาธุรกิจและ SME", "Jarvis System"},
	"L7":  {"L7", "ORACLE", "พลายแก้ว", "สถาปนิกระบบความเร็วแสง", "Jarvis System"}, // ปรับปรุงแล้ว
	"L8":  {"L8", "GUARDIAN", "Glock", "กล่องดำผู้พิทักษ์เอกสารกฎหมาย", "Jarvis System"},
	"L9":  {"L9", "SENTINEL", "Sentinel", "ผู้ป้องกันการรั่วไหลของ AI (Era of Fire)", "Jarvis System"},
	"L10": {"L10", "AUDITOR", "Auditor", "ผู้ตรวจสอบความเสี่ยงและกฎหมายสัจจะ", "Jarvis System"},
	"L11": {"L11", "BALANCER", "Balancer", "ผู้บาลานซ์ระบบและ Self-Healing", "Jarvis System"},
}
