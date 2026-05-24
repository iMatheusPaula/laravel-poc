package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func startSubscriber(config Config) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Pubsub Client error: %v", err)
	}

	defer client.Close()

	sub := client.Subscription(config.SubscriptionID)

	log.Printf("Subscribing to topic: %s", config.SubscriptionID)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handleMessage(msg.Data)
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Error receiving message: %v", err)
	}
}
