package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, helloHandler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Say /hello",
	})
}
