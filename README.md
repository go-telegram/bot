# Golang Telegram Bot

[![Go Report Card](https://goreportcard.com/badge/github.com/go-telegram/bot)](https://goreportcard.com/report/github.com/go-telegram/bot) [![codecov](https://codecov.io/gh/go-telegram/bot/branch/main/graph/badge.svg?token=57B1OR6PCK)](https://codecov.io/gh/go-telegram/bot)

✅ Present in the list of libraries https://core.telegram.org/bots/samples#go

> [Telegram Group](https://t.me/gotelegrambotui)

> Supports Bot API version: [9.0](https://core.telegram.org/bots/api#april-11-2025) from April 11, 2025

It's a Go zero-dependencies telegram bot framework

A simple example `echo-bot`:

```go
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
```

You can find more examples in the [examples](examples) folder.

To run the examples, set the `EXAMPLE_TELEGRAM_BOT_TOKEN` environment variable to your bot token.

## Getting started

Go version: **1.18**

Install the dependencies:

```bash
go get -u github.com/go-telegram/bot
```

Initialize and run the bot:

```go
b, err := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER")

b.Start(context.TODO())
```

On create bot will call the `getMe` method (with 5 sec timeout). And returns error on fail.
If you want to change this timeout, use option `bot.WithCheckInitTimeout`

You can define a default handler for the bot:

```go
b, err := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER", bot.WithDefaultHandler(handler))

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// this handler will be called for all updates
}
```

## Webhooks

If you want to use webhooks, instead of using `bot.Start`, you should use the `bot.StartWebhook` method to start the bot.
Also, you should use `bot.WebhookHandler()` method as HTTP handler for your server.

```go
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithWebhookSecretToken(os.Getenv("EXAMPLE_TELEGRAM_WEBHOOK_SECRET_TOKEN"))
	}

	b, _ := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

	// call methods.SetWebhook if needed

	go b.StartWebhook(ctx)

	http.ListenAndServe(":2000", b.WebhookHandler())

	// call methods.DeleteWebhook if needed
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
```

[Demo in examples](examples/echo_with_webhook/main.go)

Also, you can manually process updates with `bot.ProcessUpdate` method.

```go
update := models.Update{}

json.NewDecoder(req.Body).Decode(&update)

b.ProcessUpdate(ctx, &update)
```

## Middlewares

You can use middlewares with `WithMiddlewares(middlewares ...Middleware)` option.

See an example in [examples](examples/middleware/main.go)

## Available methods

All available methods are listed in the [Telegram Bot API documentation](https://core.telegram.org/bots/api)

You can use all these methods as bot funcs. All methods have name like in official documentation, but with capital first letter.

`bot.SendMessage`, `bot.GetMe`, `bot.SendPhoto`, etc

All methods have signature `(ctx context.Context, params <PARAMS>) (<response>, error)`.
Except `GetMe`, `Close` and `Logout` which are have not params

`<PARAMS>` is a struct with fields that corresponds to Telegram Bot API parameters.
All Params structs have name like for corresponded methods, but with `Params` suffix.

`SendMessageParams` for `SendMessage` method etc.

You should pass params by pointer

```go
bot.SendMessage(ctx, &bot.SendMessageParams{...})
```

## Options

You can use options to customize the bot.

```go
b, err := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER", opts...)
```

### Options list (see [options.go](options.go) for more details)

- `WithCheckInitTimeout(timeout time.Duration)` - timeout for check init bot
- `WithMiddlewares(middlewares ...Middleware)` - add middlewares
- `WithMessageTextHandler(pattern string, matchType MatchType, handler HandlerFunc)` - add handler for Message.Text field
- `WithCallbackQueryDataHandler(pattern string, matchType MatchType, handler HandlerFunc)` - add handler for CallbackQuery.Data field
- `WithPhotoCaptionHandler` - add handler for Message.Caption field
- `WithDefaultHandler(handler HandlerFunc)` - add default handler
- `WithDebug()` - enable debug mode
- `WithErrorsHandler(handler ErrorsHandler)` - add errors handler
- `WithDebugHandler(handler DebugHandler)` - add debug handler
- `WithHTTPClient(pollTimeout time.Duration, client HttpClient)` - set custom http client
- `WithServerURL(serverURL string)` - set server url
- `WithSkipGetMe()` - skip call GetMe on bot init
- `WithAllowedUpdates(params AllowedUpdates)` - set [allowed_updates](https://core.telegram.org/bots/api#getupdates) for getUpdates method
- `WithUpdatesChannelCap(cap int)` - set updates channel capacity, by default 1024
- `WithWebhookSecretToken(webhookSecretToken string)` - set X-Telegram-Bot-Api-Secret-Token header sent from telegram servers to confirm validity of update
- `WithWorkers` - set the number of workers that are processing the Updates channel, by default 1
- `UseTestEnvironment()` - use test environment
- `WithNotAsyncHandlers()` - allows to run handlers in the main goroutine
- `WithInitialOffset(offset int64)` - allows to set initial offset for getUpdates method

## Message.Text and CallbackQuery.Data handlers

For your convenience, you can use `Message.Text`, `CallbackQuery.Data` and `Message.Caption` handlers.

An example:

```go
b, err := bot.New("YOUR_BOT_TOKEN_FROM_BOTFATHER")

b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, myStartHandler)

b.Start(context.TODO())
```

> also you can use bot init options `WithMessageTextHandler` and `WithCallbackQueryDataHandler`


In this example, the handler will be called when the user sends `/start` message. All other messages will be handled by the default handler.

Handler Types:
- `HandlerTypeMessageText` - for Update.Message.Text field
- `HandlerTypeCallbackQueryData` - for Update.CallbackQuery.Data field
- `HandlerTypeCallbackQueryGameShortName` - for Update.CallbackQuery.GameShortName field
- `HandlerTypePhotoCaption` - for Update.Message.Caption field

RegisterHandler returns a handler ID string. You can use it to remove the handler later.

```
b.UnregisterHandler(handlerID)
```

Match Types:
- `MatchTypeExact`
- `MatchTypePrefix`
- `MatchTypeContains`
- `MatchTypeCommand`
- `MatchTypeCommandStartOnly`

> For `MatchTypeCommand` and `MatchTypeCommandStartOnly` usage see an [example](examples/command_handler/main.go)

You can use `RegisterHandlerRegexp` to match by regular expression.

```go
re := regexp.MustCompile(`^/start`)

b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, re, myStartHandler)
```

If you want to use custom handler, use `RegisterHandlerMatchFunc`

```go
matchFunc := func(update *models.Update) bool {
	// your checks
	return true
}

b.RegisterHandlerMatchFunc(bot.HandlerTypeMessageText, matchFunc, myHandler)
```

## InputFile

For some methods, like `SendPhoto`, `SendAudio` etc, you can send file by file path or file contents.

To send a file by URL or FileID, you can use `&models.InputFileString{Data: string}`:

```go
// file id of uploaded image
inputFileData := "AgACAgIAAxkDAAIBOWJimnCJHQJiJ4P3aasQCPNyo6mlAALDuzEbcD0YSxzjB-vmkZ6BAQADAgADbQADJAQ"
// or URL image path
// inputFileData := "https://example.com/image.png"

params := &bot.SendPhotoParams{
    ChatID:  chatID,
    Photo:   &models.InputFileString{Data: inputFileData},
}

bot.SendPhoto(ctx, params)
```

[Demo in examples](examples/send_photo/main.go)

To send an image file by its contents, you can use `&models.InputFileUpload{Filename: string, Data: io.Reader}`:

```go
fileContent, _ := os.ReadFile("/path/to/image.png")

params := &bot.SendPhotoParams{
    ChatID:  chatID,
    Photo:   &models.InputFileUpload{Filename: "image.png", Data: bytes.NewReader(fileContent)},
}

bot.SendPhoto(ctx, params)
```

[Demo in examples](examples/send_photo_upload/main.go)

## InputMedia

For methods like `SendMediaGroup` or `EditMessageMedia` you can send media by file path or file contents.

[Official documentation InputMedia](https://core.telegram.org/bots/api#inputmedia)

> field `media`: File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.

If you want to use `attach://` format, you should to define `MediaAttachment` field with file content reader.

```go
fileContent, _ := os.ReadFile("/path/to/image.png")

media1 := &models.InputMediaPhoto{
	Media: "https://telegram.org/img/t_logo.png",
}

media2 := &models.InputMediaPhoto{
	Media: "attach://image.png",
	Caption: "2",
	MediaAttachment: bytes.NewReader(fileContent),
}

params := &bot.SendMediaGroupParams{
    ChatID: update.Message.Chat.ID,
    Media: []models.InputMedia{
        media1,
        media2,
    },
}

bot.SendMediaGroup(ctx, params)
```

[Demo in examples](examples/send_media_group/main.go)

## Helpers

### `EscapeMarkdown(s string) string`

Escape special symbols for Telegram MarkdownV2 syntax

### `EscapeMarkdownUnescaped(s string) string`

Escape only unescaped special symbols for Telegram MarkdownV2 syntax

### `RandomString(n int) string`

Returns fast random a-zA-Z string with n length

### `True() bool`, `False() bool`

Allows you to define `*bool` values for params, which require `*bool`, like `SendPollParams`

```go
p := &bot.SendPollParams{
    ChatID: chatID,
    Question: "Question",
    Options: []string{"Option 1", "Option 2"},
    IsAnonymous: bot.False(),
}

b.SendPoll(ctx, p)
```

### `ValidateWebappRequest(values url.Values, token string) (user *WebAppUser, ok bool)`

Validate request from Telegram Webapp

https://core.telegram.org/bots/webapps#validating-data-received-via-the-mini-app

```go
// get url values from request
values := req.URL.Query()

user, ok := bot.ValidateWebappRequest(values, os.Getenv("TELEGRAM_BOT_TOKEN"))
if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
}
```

### `FileDownloadLink(f *models.File) string`

Return file download link after call method `GetFile`

See [documentation](https://core.telegram.org/bots/api#getfile)

## Errors

This library includes error handling. It provides the following error types:

- **ErrorForbidden (403):** This error occurs when the bot has no access to the action, such as when the user has blocked the bot.
- **ErrorBadRequest (400):** This error indicates a bad request made to the bot's API.
- **ErrorUnauthorized (401):** This error occurs when the bot's access is unauthorized for the requested action.
- **TooManyRequestsError: (429)** This error indicates that the bot has received too many requests within a short period. It includes a RetryAfter value indicating when to retry the request.
- **ErrorNotFound (404):** This error indicates that the requested resource was not found.
- **ErrorConflict (409):** This error indicates a conflict occurred during the request.

Usage:
```go
_, err := b.SendMessage(...)

if errors.Is(err, mybot.ErrorForbidden) {
    // Handle the ErrorForbidden (403) case here
}

if errors.Is(err, mybot.ErrorBadRequest) {
    // Handle the ErrorBadRequest (400) case here
}

if errors.Is(err, mybot.ErrorUnauthorized) {
    // Handle the ErrorUnauthorized (401) case here
}

if mybot.IsTooManyRequestsError(err) {
    // Handle the TooManyRequestsError (429) case here
    fmt.Println("Received TooManyRequestsError with retry_after:", err.(*mybot.TooManyRequestsError).RetryAfter)
}

if errors.Is(err, mybot.ErrorNotFound) {
    // Handle the ErrorNotFound (404) case here
}

if errors.Is(err, mybot.ErrorConflict) {
    // Handle the ErrorConflict (409) case here
}
```

## Other

- `bot.ID() int64` - returns bot ID. Bot ID is a unique identifier for the bot, obtained from the token as first part before `:`. Example: `110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw` - bot ID is `110201543`. If the bot token is invalid, the bot ID will be 0.
- `bot.Token() string` - returns bot token
- `bot.SetToken()` - set new bot token

## MiniApp misc

Check the repo [go-telegram/miniapp](https://github.com/go-telegram/miniapp) for Telegram MiniApp example.

Repo [go-telegram/miniapp-types](https://github.com/go-telegram/miniapp-types) contains TypeScript types definitions for Telegram MiniApp object.

## UI Components

In the repo https://github.com/go-telegram/ui you can find a some UI elements for your bot.

- datepicker
- inline_keyboard
- slider
- paginator

and more...

Please, check the repo for more information and live demo.
