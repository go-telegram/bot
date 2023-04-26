package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Use inline mode @botname some_text

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.InlineQuery == nil {
		return
	}

	results := []models.InlineQueryResult{
		&models.InlineQueryResultArticle{ID: "1", Title: "Foo 1", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 1"}},
		&models.InlineQueryResultArticle{ID: "2", Title: "Foo 2", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 2"}},
		&models.InlineQueryResultArticle{ID: "3", Title: "Foo 3", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 3"}},
	}

	b.AnswerInlineQuery(ctx, &bot.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})
}
