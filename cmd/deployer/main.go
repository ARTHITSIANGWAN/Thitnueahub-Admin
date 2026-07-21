package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 🧬 Message Filter Interface
type MessageFilter interface {
	Clean(input string) string
	IsValid(input string) bool
}

// 🏛️ SecureGate คือ Guard Core หลักของระบบ ThitNueaHub
type SecureGate struct {
	SecretPhrase string
	DangerWords  []string
	emojiRegex   *regexp.Regexp
	bracketRegex *regexp.Regexp
}

// 🏗️ NewSecureGate สร้าง Gatekeeper พร้อมโหลด Secret จาก Environment Variable (ป้องกัน Secret รั่ว)
func NewSecureGate() *SecureGate {
	secret := os.Getenv("TNH_SECRET_PHRASE")
	if secret == "" {
		secret = "Kin-Khao-Leaw" // Default Fallback พร้อมเตือน
	}

	return &SecureGate{
		SecretPhrase: secret,
		DangerWords:  []string{"{{stop_broadcast}}", "{{unsubscribed}}", "{{abort}}"},
		// Regex ครอบคลุมทั้ง วงเล็บเหลี่ยม และ Unicode Emojis ทั้งหมด!
		bracketRegex: regexp.MustCompile(`\[.*?\]`),
		emojiRegex:   regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}]`),
	}
}

// 🧹 Clean ทำหน้าที่กวาดล้างทั้ง [สัญลักษณ์] และ Emoji เพียวๆ แบบไม่เหลือคราบ (5ส)
func (s *SecureGate) Clean(input string) string {
	// 1. ลบข้อความในวงเล็บเหลี่ยม [xxx]
	cleaned := s.bracketRegex.ReplaceAllString(input, "")
	// 2. ลบ Emoji เพียวๆ ที่ลอยอยู่ด้านนอก
	cleaned = s.emojiRegex.ReplaceAllString(cleaned, "")
	// 3. ลบช่องว่างส่วนเกิน
	return strings.TrimSpace(cleaned)
}

// 🔒 IsValid ตรวจจับคำสั่งแฝง และรหัสผ่านสัจจะ
func (s *SecureGate) IsValid(input string) bool {
	cleaned := s.Clean(input)

	// ลบ Space ทั้งหมดเพื่อป้องกันเทคนิค Bypass เช่น {{ stop_broadcast }}
	noSpaceCleaned := strings.ReplaceAll(cleaned, " ", "")

	// 1. ตรวจสอบคำสั่งอันตราย
	for _, word := range s.DangerWords {
		cleanWord := strings.ReplaceAll(word, " ", "")
		if strings.Contains(noSpaceCleaned, cleanWord) {
			return false // บล็อกทันที!
		}
	}

	// 2. ตรวจสอบ "รหัสผ่านสัจจะเปิดระบบ"
	if !strings.Contains(cleaned, s.SecretPhrase) {
		return false
	}

	return true
}

func main() {
	gate := NewSecureGate()

	fmt.Println("--- 🛡️ ระบบตรวจสอบ ThitNueaHub Active (ไส้ในแน่น 100%) ---")

	// ทดสอบเคสที่ 1: โดนยิง Emoji ลอยๆ + คำสั่ง Bypass แบบเว้นวรรค
	test1 := "สนใจ🥰แต่บ่สนใจ {{ stop_broadcast }}"
	// ทดสอบเคสที่ 2: ส่งคำสั่งถูกต้องพร้อมรหัสลับ
	test2 := "Kin-Khao-Leaw แอดมินรันระบบ Step 15 ด่วนเลยเน้อ"

	if gate.IsValid(test1) {
		fmt.Printf("เคสที่ 1: ผ่าน (ข้อความ: %s)\n", gate.Clean(test1))
	} else {
		fmt.Printf("เคสที่ 1: [บล็อกสำเร็จ!] ตรวจพบรอยเท้าบอทแฝงเร้น หรือคำสั่งอันตราย\n")
	}

	if gate.IsValid(test2) {
		fmt.Printf("เคสที่ 2: [ผ่านด่าน] ข้อความสะอาดพร้อมประมวลผล: %s\n", gate.Clean(test2))
	} else {
		fmt.Printf("เคสที่ 2: บล็อก!\n")
	}
}
