package bot

import (
	"context"
	"sync"
	"testing"

	"github.com/go-telegram/bot/models"
	"github.com/stretchr/testify/assert"
)

// Define the necessary types for tests
type MockMiddleware struct{}

func (m MockMiddleware) Handle(next HandlerFunc) HandlerFunc {
	return func(ctx context.Context, b *Bot, upd *models.Update) {
		next(ctx, b, upd)
	}
}

type MockHandler struct{}

func (h MockHandler) Handle(ctx context.Context, b *Bot, upd *models.Update) {}

func TestApplyMiddlewares(t *testing.T) {
	finalHandlerCalled := false
	finalHandler := func(ctx context.Context, b *Bot, upd *models.Update) {
		finalHandlerCalled = true
	}

	middleware1Called := false
	middleware1 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, b *Bot, upd *models.Update) {
			middleware1Called = true
			next(ctx, b, upd)
		}
	}

	middleware2Called := false
	middleware2 := func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context, b *Bot, upd *models.Update) {
			middleware2Called = true
			next(ctx, b, upd)
		}
	}

	handler := applyMiddlewares(finalHandler, middleware1, middleware2)
	handler(context.Background(), nil, nil)

	assert.True(t, middleware1Called, "Expected middleware1 to be called")
	assert.True(t, middleware2Called, "Expected middleware2 to be called")
	assert.True(t, finalHandlerCalled, "Expected final handler to be called")
}

func TestProcessUpdate(t *testing.T) {
	b := &Bot{
		defaultHandlerFunc: func(ctx context.Context, b *Bot, upd *models.Update) {},
		handlers:           make(map[string]handler),
		middlewares:        []Middleware{},
		handlersMx:         &sync.RWMutex{},
	}

	b.defaultHandlerFunc = func(ctx context.Context, b *Bot, upd *models.Update) {
		assert.NotNil(t, upd, "Expected update to be processed")
	}

	upd := &models.Update{Message: &models.Message{}}
	b.ProcessUpdate(context.Background(), upd)
}

func TestProcessUpdateWithMessage(t *testing.T) {
	b := &Bot{
		defaultHandlerFunc: func(ctx context.Context, b *Bot, upd *models.Update) {},
		handlers:           make(map[string]handler),
		middlewares:        []Middleware{},
		handlersMx:         &sync.RWMutex{},
	}

	handlerCalled := false
	b.RegisterHandler(HandlerTypeMessageText, "pattern", MatchTypeExact, func(ctx context.Context, b *Bot, upd *models.Update) {
		handlerCalled = true
	})

	upd := &models.Update{Message: &models.Message{Text: "pattern"}}
	b.ProcessUpdate(context.Background(), upd)

	assert.True(t, handlerCalled, "Expected message handler to be called")
}

func TestProcessUpdateWithCallbackQuery(t *testing.T) {
	b := &Bot{
		defaultHandlerFunc: func(ctx context.Context, b *Bot, upd *models.Update) {},
		handlers:           make(map[string]handler),
		middlewares:        []Middleware{},
		handlersMx:         &sync.RWMutex{},
	}

	handlerCalled := false
	b.RegisterHandler(HandlerTypeCallbackQueryData, "pattern", MatchTypeExact, func(ctx context.Context, b *Bot, upd *models.Update) {
		handlerCalled = true
	})

	upd := &models.Update{CallbackQuery: &models.CallbackQuery{Data: "pattern"}}
	b.ProcessUpdate(context.Background(), upd)

	assert.True(t, handlerCalled, "Expected callback query handler to be called")
}

func TestFindHandler(t *testing.T) {
	b := &Bot{
		defaultHandlerFunc: func(ctx context.Context, b *Bot, upd *models.Update) {},
		handlers:           make(map[string]handler),
		handlersMx:         &sync.RWMutex{},
	}

	handlerCalled := false
	b.RegisterHandler(HandlerTypeMessageText, "pattern", MatchTypeExact, func(ctx context.Context, b *Bot, upd *models.Update) {
		handlerCalled = true
	})

	upd := &models.Update{Message: &models.Message{Text: "pattern"}}
	foundHandler := b.findHandler(HandlerTypeMessageText, upd)

	assert.NotNil(t, foundHandler, "Expected handler to be found")
	foundHandler(context.Background(), b, upd)
	assert.True(t, handlerCalled, "Expected found handler to be called")
}
