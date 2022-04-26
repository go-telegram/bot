# Golang Telegram Bot

> The project is under development. API may be changed before v1.0.0 version.
> Not all methods are implemented yet.

It's a Go zero-dependencies telegram bot framework

An `echo-bot` simple example:

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER", opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: strconv.Itoa(update.Message.Chat.ID),
		Text:   update.Message.Text,
	})
}
```

You can find more examples in the examples folder.

For test examples, you should set environment variable `EXAMPLE_TELEGRAM_BOT_TOKEN` to your bot token.

## Getting started

Go version: **1.18**

Install the dependencies:

```bash
go get -u github.com/go-telegram/bot
```

Initialize and run the bot:

```go
b := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER")

b.Start(context.TODO())
```

## Available methods

All available methods are listed in the [Telegram Bot API documentation](https://core.telegram.org/bots/api)

You can find these methods in the `methods` package. All methods have name like in official documentation, but with capital first letter.

`methods.SendMessage`, `methods.GetMe`, `methods.SendPhoto`, etc

## Options

You can use options to customize the bot.

Full list of options you can find [here](options.go)

## Message.Text and CallbackQuery.Data handlers

For your convenience, you can use `Message.Text` and `CallbackQuery.Data` handlers.

An example:

```go

b := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER")

b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, myStartHandler)

b.Start(context.TODO())
```

In this example, the handler will be called when the user sends `/start` message. All other messages will be handled by the default handler.

Handler Types:
- `HandlerTypeMessageText` - for Update.Message.Text field
- `HandlerTypeCallbackQueryData` - for Update.CallbackQuery.Data field

Match Types:
- `MatchTypeExact` 
- `MatchTypePrefix` 
- `MatchTypeContains`
- `MatchTypeRegexp`

## UI

In the repo https://github.com/go-telegram/ui you can find a some UI elements for your bot.

- datepicker
- inline_keyboard
- slider
- paginator

and more...

Please, check the repo for more information.
