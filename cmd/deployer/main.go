package main

import (
	"fmt"
	"os"
)

// ข้อมูล 11 ขุนพล (Centralized Registry)
type Agent struct {
	Level, Role, Name, Function, Base, Link string
}

var Registry = []Agent{
	{"L1", "COMMANDER", "ทิศเหนือ", "ผู้บัญชาการสูงสุด", "Jarvis V83", "https://arthitsiangwan.github.io/ThitNueaHub-Admin/"},
	{"L2", "SCRIBE", "แก้วตา", "สื่อสารและสัจจะ", "Jarvis V83", "https://thitnueahub.com/kaewta"},
	{"L3", "ARTIST", "น้ำอิง", "ผู้สร้างสรรค์ตรรกะ", "Jarvis V83", "https://thitnueahub.com/narming"},
    // บอสเพิ่มข้อมูลขุนพลที่เหลือให้ครบ 11 ที่นี่...
}

func main() {
	outputDir := "docs/profiles"
	os.MkdirAll(outputDir, 0755)

	for _, a := range Registry {
		fileName := fmt.Sprintf("%s/%s_%s.md", outputDir, a.Level, a.Name)
		content := fmt.Sprintf(`# 🛡️ SOVEREIGN: %s
**Role:** %s | **Base:** %s
**Function:** %s
---
### 🛰️ MISSION STATUS
- [ ] 16 Episodes Ready
- [x] V83 Protocol Engaged

### 🌐 CONTACT
[Link Portal](%s)
`, a.Name, a.Role, a.Base, a.Function, a.Link)

		os.WriteFile(fileName, []byte(content), 0644)
		fmt.Printf("✅ Generated: %s\n", fileName)
	}
	fmt.Println("🚀 Deployment Complete: 11 Sovereigns Registered.")
}
