package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
)

func main() {
	b := bot.New(context.Background(), os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))

	user, err := methods.GetMe(context.Background(), b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User: %#v\n", user)
}
