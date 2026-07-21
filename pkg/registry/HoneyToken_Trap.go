package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"os"
	"time"
)

// HoneyTrapLog โครงสร้างข้อมูลสำหรับส่ง Alert หาบอส
type HoneyTrapLog struct {
	Timestamp string `json:"timestamp"`
	ClientIP  string `json:"client_ip"`
	UserAgent string `json:"user_agent"`
	URI       string `json:"requested_uri"`
	Method    string `json:"method"`
	Payload   string `json:"payload,omitempty"`
}

// TelegramWebhookURL รับค่าจาก Environment Variable (ซ่อนไว้หลังบ้าน)
var telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
var telegramChatID = os.Getenv("TELEGRAM_CHAT_ID")

func main() {
	// Endpoint ล่อซื้อ (Decoy Routes)
	http.HandleFunc("/api/v1/admin/config", honeyTrapHandler)
	http.HandleFunc("/v1/auth/secret-key", honeyTrapHandler)
	http.HandleFunc("/.env", honeyTrapHandler)

	port := ":8080"
	fmt.Printf("[+] HoneyToken Trap Active on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server Error: %v\n", err)
	}
}

func honeyTrapHandler(w http.ResponseWriter, r *http.Request) {
	// ดักจับ IP
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}

	logData := HoneyTrapLog{
		Timestamp: time.Now().Format(time.RFC3339),
		ClientIP:  ip,
		UserAgent: r.UserAgent(),
		URI:       r.RequestURI,
		Method:    r.Method,
	}

	// ส่งแจ้งเตือนฉุกเฉินหาบอสทันที
	go sendAlertToBoss(logData)

	// ตอบกลับด้วยข้อความหลอก (Decoy Response) ให้ฝั่งนู้นงงว่าทำไมใช้ไม่ได้
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"status": "error", "code": 401, "message": "Invalid Auth Token or Expired Session"}`))
}

func sendAlertToBoss(log HoneyTrapLog) {
	if telegramBotToken == "" || telegramChatID == "" {
		return
	}

	msg := fmt.Sprintf("🚨 *[INTRUDER ALERT] มีผู้หลงเข้ามาในค่ายกล!*\n"+
		"⏰ **เวลา:** `%s`\n"+
		"🌐 **IP:** `%s`\n"+
		"🎯 **Target URI:** `%s`\n"+
		"🔍 **User-Agent:** `%s`",
		log.Timestamp, log.ClientIP, log.URI, log.UserAgent)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	body, _ := json.Marshal(map[string]string{
		"chat_id":    telegramChatID,
		"text":       msg,
		"parse_mode": "Markdown",
	})

	http.Post(url, "application/json", bytes.NewBuffer(body))
}
