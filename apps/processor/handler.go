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

func handleMessage(data json.RawMessage) {
	var msg AppointmentMessage

	err := json.Unmarshal(data, &msg)

	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
	}

	sendEmail(msg)
}

func sendEmail(msg AppointmentMessage) {
	log.Println("==============================")
	log.Println("📧 Simulando envio de email...")
	log.Printf("   Para: %s <%s>", msg.ContactName, msg.ContactEmail)
	log.Printf("   Assunto: Confirmação de Appointment #%d", msg.AppointmentId)
	log.Printf("   Agendado para: %s", msg.ScheduledAt)
	log.Println("   Status: Enviado com sucesso!")
	log.Println("==============================")
}
