package bot

import (
	"context"
	"testing"

	"github.com/go-telegram/bot/models"
)

func Test_applyMiddlewares(t *testing.T) {
	h := func(ctx context.Context, bot *Bot, update *models.Update) {}

	middleware1 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, bot *Bot, update *models.Update) {
			next(ctx, bot, update)
		}
	}

	middleware2 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, bot *Bot, update *models.Update) {
			next(ctx, bot, update)
		}
	}

	wrapped := applyMiddlewares(h, middleware1, middleware2)
	if wrapped == nil {
		t.Fatal("Expected wrapped handler to be non-nil")
	}
}

func TestProcessUpdate(t *testing.T) {
	var called bool
	h := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = true
	}

	bot := &Bot{
		defaultHandlerFunc: h,
		middlewares:        []Middleware{},
		notAsyncHandlers:   true,
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	bot.ProcessUpdate(ctx, upd)
	if !called {
		t.Fatal("Expected default handler to be called")
	}
}

func TestProcessUpdate_WithMiddlewares(t *testing.T) {
	var called bool
	h := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = true
	}

	middleware := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, bot *Bot, update *models.Update) {
			next(ctx, bot, update)
		}
	}

	bot := &Bot{
		defaultHandlerFunc: h,
		middlewares:        []Middleware{middleware},
		notAsyncHandlers:   true,
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	bot.ProcessUpdate(ctx, upd)
	if !called {
		t.Fatal("Expected default handler to be called")
	}
}

func TestProcessUpdate_WithMatchTypeFunc(t *testing.T) {
	var called string
	h1 := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = "h1"
	}
	h2 := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = "h2"
	}
	m := func(update *models.Update) bool {
		return update.CallbackQuery.GameShortName == "game"
	}

	bot := &Bot{
		defaultHandlerFunc: h1,
		notAsyncHandlers:   true,
	}

	bot.RegisterHandlerMatchFunc(m, h2)

	ctx := context.Background()
	upd := &models.Update{ID: 42, CallbackQuery: &models.CallbackQuery{ID: "test", GameShortName: "game"}}

	bot.ProcessUpdate(ctx, upd)
	if called != "h2" {
		t.Fatalf("Expected h2 handler to be called but %s handler was called", called)
	}
}

func Test_findHandler(t *testing.T) {
	var called bool
	h := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = true
	}

	bot := &Bot{
		defaultHandlerFunc: h,
	}

	// Register a handler
	bot.handlers = append(bot.handlers, handler{
		id:          "test",
		handlerType: HandlerTypeMessageText,
		matchType:   MatchTypeExact,
		pattern:     "test",
		handler:     h,
	})

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	handler := bot.findHandler(upd)
	handler(ctx, bot, upd)

	if !called {
		t.Fatal("Expected registered handler to be called")
	}
}

func Test_findHandler_Default(t *testing.T) {
	var called bool
	h := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = true
	}

	bot := &Bot{
		defaultHandlerFunc: h,
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	handler := bot.findHandler(upd)
	handler(ctx, bot, upd)

	if !called {
		t.Fatal("Expected default handler to be called")
	}
}
