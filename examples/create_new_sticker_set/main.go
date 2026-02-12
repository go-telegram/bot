package main

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/TGlimmer/gobot"
	"github.com/TGlimmer/gobot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithSkipGetMe(),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	user, err := b.GetMe(ctx)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
	botUsername = user.Username

	b.Start(ctx)
}

//go:embed images
var images embed.FS
var botUsername string

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fileContent, _ := images.ReadFile("images/telegram.png")

	inputSticker1 := models.InputSticker{
		Sticker:   "https://github.com/TGlimmer/gobot/blob/main/examples/create_new_sticker_set/images/telegram.png?raw=true",
		Format:    "static",
		EmojiList: []string{"1️⃣"},
	}

	inputSticker2 := models.InputSticker{
		Sticker:           "attach://telegram.png",
		Format:            "static",
		EmojiList:         []string{"2️⃣"},
		StickerAttachment: bytes.NewReader(fileContent),
	}

	stickerSetName := fmt.Sprintf("Example%d_by_%s", time.Now().Unix(), botUsername)
	params := &bot.CreateNewStickerSetParams{
		UserID: update.Message.Chat.ID,
		Name:   stickerSetName,
		Title:  "Example sticker set",
		Stickers: []models.InputSticker{
			inputSticker1,
			inputSticker2,
		},
	}

	_, err := b.CreateNewStickerSet(ctx, params)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	stickerSet, err := b.GetStickerSet(ctx, &bot.GetStickerSetParams{Name: stickerSetName})
	if err != nil {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
	_, err = b.SendSticker(
		ctx,
		&bot.SendStickerParams{
			ChatID: update.Message.Chat.ID,
			Sticker: &models.InputFileString{
				Data: stickerSet.Stickers[0].FileID,
			},
		},
	)
	if err != nil {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
}
