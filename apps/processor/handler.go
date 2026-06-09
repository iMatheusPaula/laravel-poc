package main

import (
	"encoding/json"
	"log"
)

type AppointmentMessage struct {
	AppointmentId int    `json:"appointment_id"`
	ContactName   string `json:"contact_name"`
	ContactEmail  string `json:"contact_email"`
	ScheduledAt   string `json:"scheduled_at"`
}

func handleMessage(data json.RawMessage, mailer *Mailer) bool {
	var msg AppointmentMessage

	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return false
	}

	return mailer.Send(msg)
}
