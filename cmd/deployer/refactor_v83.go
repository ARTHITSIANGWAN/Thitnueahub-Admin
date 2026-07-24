package main

import (
	"fmt"
	"os"
	"strings"
)

// ฟังก์ชันสำหรับจำลองและตรวจสอบการแทนที่ค่าเวอร์ชันเก่าในโปรเจกต์
func RefactorVersionString(content string, oldStr string, newStr string) (string, int) {
	// นับจำนวนจุดที่พบคำว่า v83 เก่า
	count := strings.Count(content, oldStr)
	// แทนที่ด้วยชื่อใหม่ที่ถูกต้อง
	updatedContent := strings.ReplaceAll(content, oldStr, newStr)
	return updatedContent, count
}

func main() {
	// ตัวอย่างข้อมูลโค้ดจำลองที่ยังติดค่า v83
	sampleCodeConfig := `
		AppIdentifier = "thitnueahub-admin-v83"
		ActiveVersion = "v83"
		RoutingPath = "/api/v83/workers"
	`

	oldTarget := "v83"
	newTarget := "v84-cluster" // เปลี่ยนเป็นชื่อเวอร์ชันหรือชื่อใหม่ที่ต้องการ

	fmt.Println("=== เริ่มต้นกระบวนการตรวจสอบค่าเวอร์ชันเก่าในโค้ด ===")
	
	result, replacedCount := RefactorVersionString(sampleCodeConfig, oldTarget, newTarget)
	
	fmt.Printf("พบคำว่า '%s' ทั้งหมด: %d จุด\n", oldTarget, replacedCount)
	fmt.Println("--- ผลลัพธ์หลังอัปเดตโค้ด ---")
	fmt.Println(result)

	// หากต้องการบันทึกจริง สามารถเขียนทับไฟล์โปรเจกต์ได้ผ่าน os.WriteFile
	_ = os.WriteFile("config_updated.txt", []byte(result), 0644)
	fmt.Println("=== บันทึกการอัปเดตโครงสร้างสำเร็จ ===")
}

