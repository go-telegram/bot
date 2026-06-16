package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Inline rich message example.
//
// Use inline mode (@botname some_text) to trigger the bot. The bot returns a
// single InlineQueryResultArticle candidate. When the user selects it, the bot
// sends a Bot API 10.1 Rich Message that contains a title, a description and a
// photo block.
//
// Notes about the image URL:
//   - thumbnail_url is only used for the inline candidate list preview.
//   - The Rich Message photo URL is provided inside the markdown media block and
//     is NOT shown as plain text in the message body.
//   - The photo URL is still fetched by Telegram, so it must be treated as a
//     publicly accessible resource.
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	// The token is read from the environment; never hard-code a bot token.
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

	// thumbnail_url is used for the inline candidate list only.
	const thumbnailURL = "https://cdn.example.com/thumb.jpg"
	// photoURL is a publicly accessible resource; it is embedded in the rich
	// message media block, not written into the visible message text.
	const photoURL = "https://cdn.example.com/photo.jpg"

	richMarkdown := "# Shared title\n\n" +
		"Some descriptive text for the rich message.\n\n" +
		"![Preview](" + photoURL + ")"

	results := []models.InlineQueryResult{
		&models.InlineQueryResultArticle{
			ID:           "share_1",
			Title:        "Title",
			Description:  "Description",
			ThumbnailURL: thumbnailURL,
			// The Article URL field is intentionally left unset.
			InputMessageContent: &models.InputRichMessageContent{
				RichMessage: models.InputRichMessage{
					Markdown: richMarkdown,
				},
			},
		},
	}

	b.AnswerInlineQuery(ctx, &bot.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})
}
