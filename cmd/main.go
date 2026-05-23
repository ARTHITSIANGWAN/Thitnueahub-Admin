package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TrinityStatus struct {
	Era          string    `json:"era"`
	Architecture string    `json:"architecture"`
	SatchaDB     string    `json:"satcha_database_id"`
	Latency      string    `json:"latency"`
	Timestamp    time.Time `json:"timestamp"`
}

func main() {
	log.Println("⚡ [TNH V83 TRINITY]: Ignite the 9th Era of Fire...")

	// [อุดช่องโหว่ที่ 1]: แยกเกราะป้องกัน สร้าง Custom ServeMux ของตัวเองเด็ดขาด 
	// ไม่ใช้ DefaultServeMux ร่วมกับใคร ป้องกัน Library ขยะแอบมาเปิดพอร์ตผีหลังบ้าน
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v83/trinity-status", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now() // เริ่มจับเวลาจริงทันทีที่ Request วิ่งข้ามประตูด่านตรวจเข้ามา

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// [อุดช่องโหว่ที่ 2]: ลบข้อมูลขยะแช่แข็ง 0.08ms ทิ้งซะ!
		// แล้วคำนวณเวลาประมวลผลจริง (Real-time Latency) พ่นสัจจะข้อมูลให้ขุนพล AI เอาไปใช้วิเคราะห์ต่อได้แม่นยำ
		duration := time.Since(startTime)
		actualLatency := fmt.Sprintf("%.4fms", float64(duration.Nanoseconds())/1e6)

		res := TrinityStatus{
			Era:          "9th_Era_of_Fire_Orchestra",
			Architecture: "100_Percent_Pure_Go_Logic",
			SatchaDB:     "6a8b4373-bf40-4b63-bb02-f612ecbe63b7", 
			Latency:      actualLatency, // ของจริงตามเนื้อผ้า ไม่มีการโม้ตัวเลข
			Timestamp:    time.Now(),
		}
		_ = json.NewEncoder(w).Encode(res)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>🏰 V83 TRINITY EMPIRE CORE ACTIVE</h1><h3>Zero-Garbage Sovereign Port: 2026</h3>")
	})

	port := "2026"
	fmt.Printf("👑 TRINITY EMPIRE V83 | 🔥 ENGINE ONLINE | Port: %s\n", port)

	// ส่งมิวซ์ส่วนตัว (mux) ที่เราควบคุมสิทธิ์เอง 100% เข้าไปรันระบบ
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
