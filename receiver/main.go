package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WebhookPayload struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload
	json.NewDecoder(r.Body).Decode(&payload)

	log.Println("Notification Received:", payload.Event)
	log.Println("Notification Data:", payload.Data)
	// handle logic
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", notificationHandler)
	fmt.Println("Listening on :8001")
	http.ListenAndServe(":8001", nil)
}
