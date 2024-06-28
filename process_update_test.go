package bot

import (
	"context"
	"sync"
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
		handlersMx:         &sync.RWMutex{},
		handlers:           map[string]handler{},
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
		handlersMx:         &sync.RWMutex{},
		handlers:           map[string]handler{},
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	bot.ProcessUpdate(ctx, upd)
	if !called {
		t.Fatal("Expected default handler to be called")
	}
}

func Test_findHandler(t *testing.T) {
	var called bool
	h := func(ctx context.Context, bot *Bot, update *models.Update) {
		called = true
	}

	bot := &Bot{
		defaultHandlerFunc: h,
		handlersMx:         &sync.RWMutex{},
		handlers:           map[string]handler{},
	}

	// Register a handler
	bot.handlers["test"] = handler{
		handlerType: HandlerTypeMessageText,
		matchType:   MatchTypeExact,
		pattern:     "test",
		handler:     h,
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	handler := bot.findHandler(HandlerTypeMessageText, upd)
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
		handlersMx:         &sync.RWMutex{},
		handlers:           map[string]handler{},
	}

	ctx := context.Background()
	upd := &models.Update{Message: &models.Message{Text: "test"}}

	handler := bot.findHandler(HandlerTypeCallbackQueryData, upd)
	handler(ctx, bot, upd)

	if !called {
		t.Fatal("Expected default handler to be called")
	}
}
