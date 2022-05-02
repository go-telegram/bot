package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b := bot.New(ctx, os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	methods.SetWebhook(ctx, b, &methods.SetWebhookParams{
		URL: "https://example.com/webhook",
	})

	http.ListenAndServe(":2000", b.WebhookHandler())

	// call methods.DeleteWebhook if needed
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
