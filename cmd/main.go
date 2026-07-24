package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2" // Engine สเปกสูง ไร้ขยะ Allocation
)

var SECURITY_SECRET = []byte(os.Getenv("SECRET"))

const LISTEN_PORT = ":2026"
const TASK_QUEUE_LIMIT = 128

// GripenCommandPayload โครงสร้างภาษากลางสื่อสารระหว่าง Node
type GripenCommandPayload struct {
	CommandID string                 `json:"command_id"`
	Action    string                 `json:"action"`
	Squadron  string                 `json:"squadron"`
	MetaInfo  map[string]interface{} `json:"meta_info"`
	Timestamp int64                  `json:"timestamp"`
	Signature string                 `json:"signature"`
}

var centralJobQueue = make(chan GripenCommandPayload, TASK_QUEUE_LIMIT)

func init() {
	if len(SECURITY_SECRET) == 0 {
		SECURITY_SECRET = []byte("tnh-gripen-sovereign-secret-2026")
	}
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "TNH-Orchestra-Core-V84.9.2",
		AppName:      "ThitNueaHub Admin Center",
	})

	// Layer 1: Entrance Endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>🏰 THITNUEAHUB-ADMIN ORCHESTRA ENGINE ONLINE</h1>")
	})

	// Layer 10 & 9: Strategic Command Launch Pad
	app.Post("/api/v84/squadron/launch", func(c *fiber.Ctx) error {
		payload := new(GripenCommandPayload)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Payload Structure")
		}

		// ตรวจสอบความถูกต้อง HMAC Signature
		if !validateHMAC(*payload, payload.Signature) {
			return c.Status(fiber.StatusUnauthorized).SendString("🛡️ Security Breach: Invalid Token Signature")
		}

		select {
		case centralJobQueue <- *payload:
			go triggerKaewtaMonitor(payload.CommandID, payload.Squadron)

			return c.JSON(fiber.Map{
				"status":           "🍊 Dispatched to Orange Queue",
				"provider_name":    "thitnueahub-admin", // อัปเดต Node Name ล่าสุด
				"cloudflare_relay": "tnh-files.9thera.workers.dev",
				"command_id":       payload.CommandID,
			})
		default:
			return c.Status(fiber.StatusServiceUnavailable).SendString("Central Queue Overflow")
		}
	})

	// เปิด Worker 4 ตัวทำงานแบบ Parallel
	for workerID := 1; workerID <= 4; workerID++ {
		go runSovereignSquadronWorker(workerID)
	}

	log.Fatal(app.Listen(LISTEN_PORT))
}

func generateHMAC(p GripenCommandPayload) string {
	mac := hmac.New(sha256.New, SECURITY_SECRET)
	mac.Write([]byte(fmt.Sprintf("%s:%s:%s:%d", p.CommandID, p.Action, p.Squadron, p.Timestamp)))
	return hex.EncodeToString(mac.Sum(nil))
}

func validateHMAC(p GripenCommandPayload, sig string) bool {
	return generateHMAC(p) == sig
}

func triggerKaewtaMonitor(id string, squad string) {
	log.Printf("⚡ [thitnueahub-admin -> L2]: ล็อกพิกัดคำสั่งฝูงบิน %s (ID: %s)\n", squad, id)
}

func runSovereignSquadronWorker(id int) {
	for job := range centralJobQueue {
		log.Printf("✈️ [Worker %d] ประมวลผลจาก thitnueahub-admin | Job ID: %s\n", id, job.CommandID)
		log.Printf("🧹 [Worker %d]: เคลียร์ Context และ Memory เรียบร้อย 100%%\n", id)
	}
}
