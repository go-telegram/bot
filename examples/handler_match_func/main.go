package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.RegisterHandlerMatchFunc(matchFunc, helloHandler)

	b.Start(ctx)
}

func matchFunc(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "hello"
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("hello handler")
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("default handler")
}
