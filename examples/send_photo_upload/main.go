package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
	"strconv"
)

// Send any text message to the bot after the bot has been started

func main() {
	opts := []bot.Option{
		bot.WithDebug(),
		bot.WithDefaultHandler(handler),
	}

	b := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fileData, errReadFile := os.ReadFile("./examples/send_photo_upload/facebook.png")
	if errReadFile != nil {
		fmt.Printf("error read file, %v\n", errReadFile)
		return
	}

	params := &methods.SendPhotoParams{
		ChatID:  strconv.Itoa(update.Message.Chat.ID),
		Photo:   &models.InputFileUpload{Filename: "facebook.png", Data: bytes.NewReader(fileData)},
		Caption: "New uploaded Facebook logo",
	}

	methods.SendPhoto(ctx, b, params)
}
