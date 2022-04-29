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
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}

	b := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   "You selected the button: " + update.CallbackQuery.Data,
	})
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Button 1", CallbackData: "button_1"},
				{Text: "Button 2", CallbackData: "button_2"},
			}, {
				{Text: "Button 3", CallbackData: "button_3"},
			},
		},
	}

	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Click by button",
		ReplyMarkup: kb,
	})
}
