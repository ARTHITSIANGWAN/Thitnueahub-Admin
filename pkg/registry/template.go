package registry

import "fmt"

func GenerateMarkdown(a AgentProfile) string {
	return fmt.Sprintf(`# 🛡️ SOVEREIGN: %s
**Role:** %s | **Base:** %s
**Function:** %s
---
### 🛰️ MISSION STATUS
- [ ] 16 Episodes Ready
- [x] V83 Protocol Engaged

### 🌐 CONTACT
[Link Portal](%s)
`, a.Name, a.Role, a.Base, a.Function, a.Link)
}
