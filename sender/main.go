package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WebhookPayload struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

func sendWebhook(url string, payload WebhookPayload) error {
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("✅ Webhook sent!")
	return nil
}

func main() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		payload := WebhookPayload{
			Event: "notification",
			Data:  fmt.Sprintf("Your task is completed at %s", t.Format(time.RFC3339)),
		}
		log.Println("Sending webhook event...")
		err := sendWebhook("http://localhost:8001/webhook", payload)
		if err != nil {
			fmt.Println("❌ Error sending webhook:", err)
		}
	}
}
