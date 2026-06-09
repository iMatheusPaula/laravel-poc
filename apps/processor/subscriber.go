package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func startSubscriber(ctx context.Context, config Config, mailer *Mailer) {
	client, err := pubsub.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Pubsub Client error: %v", err)
	}

	defer client.Close()

	sub := client.Subscription(config.SubscriptionID)

	log.Printf("Subscribing to topic: %s", config.SubscriptionID)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		result := handleMessage(msg.Data, mailer)
		if !result {
			msg.Nack()
			return
		}
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Error receiving message: %v", err)
	}
}
