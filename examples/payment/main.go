package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Payments example

// 1. Add payment method for your bot in BotFather. For example, use Test provider Paymaster.
// 2. Store provider token in environment variable EXAMPLE_TELEGRAM_PAYMENT_TOKEN.
// 3. Start the bot and send /payment command to the bot.
// 4. Click on the payment button in the chat.
// 5. Enter test card data: 5185730450000003, 01/30, 123 (see test cards in https://developer.mastercard.com/mastercard-merchant-presented-qr/documentation/server-apis/test-card-numbers/)
// 6. Click on the Pay button.

// See https://core.telegram.org/bots/payments#step-by-step-process for details

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithMessageTextHandler("/payment", bot.MatchTypeExact, handlerPaymentCommand),
	}

	b, _ := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	b.Start(ctx)
}

func handlerPaymentCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID:          update.Message.Chat.ID,
		MessageThreadID: 0,
		Title:           "Invoice Title",
		Description:     "Invoice description",
		Payload:         "xxx",
		ProviderToken:   os.Getenv("EXAMPLE_TELEGRAM_PAYMENT_TOKEN"),
		Currency:        "RUB",
		Prices: []models.LabeledPrice{
			{Label: "Price 1", Amount: 12425},
			{Label: "Price 2", Amount: 32454},
		},
	})
	if err != nil {
		fmt.Printf("error sending invoice: %s\n", err)
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.PreCheckoutQuery != nil {
		fmt.Printf("get PreCheckoutQuery for invoce payload: %s\n", update.PreCheckoutQuery.InvoicePayload)

		b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
			PreCheckoutQueryID: update.PreCheckoutQuery.ID,
			OK:                 true,
			ErrorMessage:       "",
		})
		return
	}

	if update.Message != nil {
		if update.Message.SuccessfulPayment != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      fmt.Sprintf("Payment was successful with payment payload: *%s*", update.Message.SuccessfulPayment.InvoicePayload),
				ParseMode: models.ParseModeMarkdown,
			})
			return
		}
	}
}
