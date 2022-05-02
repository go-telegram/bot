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
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDebug(),
		bot.WithDefaultHandler(handler),
	}

	b := bot.New(ctx, os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	b.GetUpdates(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	params := &methods.SendPhotoParams{
		ChatID:  update.Message.Chat.ID,
		Photo:   &models.InputFileString{Data: "AgACAgIAAxkDAAIBOWJimnCJHQJiJ4P3aasQCPNyo6mlAALDuzEbcD0YSxzjB-vmkZ6BAQADAgADbQADJAQ"},
		Caption: "Preloaded Facebook logo",
	}

	methods.SendPhoto(ctx, b, params)
}
