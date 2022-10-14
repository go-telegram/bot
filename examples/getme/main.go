package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
)

func main() {
	b, _ := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))

	user, _ := b.GetMe(context.Background())

	fmt.Printf("User: %#v\n", user)
}
