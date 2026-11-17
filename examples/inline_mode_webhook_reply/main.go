package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Use inline mode @botname some_text after the bot has been started

func main() {
	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.SetWebhook(context.Background(), &bot.SetWebhookParams{
		URL: "https://example.com/webhook",
	})

	b.RegisterHandlerInlineQuery(handler)

	// WebhookReplyHandler runs handlers itself. StartWebhook is not needed
	http.ListenAndServe(":2000", b.WebhookReplyHandler())

	// call methods.DeleteWebhook if needed
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	results := []models.InlineQueryResult{
		&models.InlineQueryResultArticle{ID: "1", Title: "Foo 1", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 1"}},
		&models.InlineQueryResultArticle{ID: "2", Title: "Foo 2", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 2"}},
		&models.InlineQueryResultArticle{ID: "3", Title: "Foo 3", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 3"}},
	}

	// Answer in the webhook response instead of making a request to the Bot API
	b.WebhookReply(ctx, "answerInlineQuery", &bot.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})
}
