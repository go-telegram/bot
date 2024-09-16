package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(showMessageWithUserID, showMessageWithUserName),
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.Printf("callback query data: %s", update.CallbackQuery.Data)
	}, singleFlight)

	b.Start(ctx)
}

func showMessageWithUserID(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("%d say: %s", update.Message.From.ID, update.Message.Text)
		}
		next(ctx, b, update)
	}
}

func showMessageWithUserName(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("%s say: %s", update.Message.From.FirstName, update.Message.Text)
		}
		next(ctx, b, update)
	}
}

// singleFlight is a middleware that ensures that only one callback query is processed at a time.
func singleFlight(next bot.HandlerFunc) bot.HandlerFunc {
	sf := sync.Map{}
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			key := update.CallbackQuery.Message.Message.ID
			if _, loaded := sf.LoadOrStore(key, struct{}{}); loaded {
				return
			}
			defer sf.Delete(key)
			next(ctx, b, update)
		}
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
