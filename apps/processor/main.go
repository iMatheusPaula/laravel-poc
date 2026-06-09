package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	config := LoadConfig()
	mailer := newMailer(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	startSubscriber(ctx, config, mailer)
}
