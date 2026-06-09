package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	config := LoadConfig()
	mailer := NewMailer(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	startSubscriber(ctx, config, mailer)
}
