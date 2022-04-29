package main

import (
	"bytes"
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
		bot.WithDefaultHandler(handler),
	}

	b := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fileDataFacebook, _ := os.ReadFile("./examples/send_media_group/facebook.png")
	fileDataYoutube, _ := os.ReadFile("./examples/send_media_group/youtube.png")

	media1 := &models.InputMediaPhoto{
		Media:   "https://telegram.org/img/t_logo.png",
		Caption: "Telegram Logo",
	}

	media2 := &models.InputMediaPhoto{
		Media:           "attach://facebook.png",
		Caption:         "Facebook Logo",
		MediaAttachment: bytes.NewReader(fileDataFacebook),
	}

	media3 := &models.InputMediaPhoto{
		Media:           "attach://youtube.png",
		Caption:         "Youtube Logo",
		MediaAttachment: bytes.NewReader(fileDataYoutube),
	}

	params := &methods.SendMediaGroupParams{
		ChatID: update.Message.Chat.ID,
		Media: []models.InputMedia{
			media1,
			media2,
			media3,
		},
	}

	methods.SendMediaGroup(ctx, b, params)
}
