package bot

import (
	"regexp"
	"sync"
	"testing"

	"github.com/go-telegram/bot/models"
)

func Test_match_func(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	var called bool

	id := b.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		called = true
		if update.ID != 42 {
			t.Error("invalid update id")
		}
		return true
	}, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{ID: 42})
	if !called {
		t.Error("not called")
	}
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_exact(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id := b.RegisterHandler(HandlerTypeMessageText, "xxx", MatchTypeExact, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{Message: &models.Message{Text: "zzz"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{Message: &models.Message{Text: "xxx"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_prefix(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id := b.RegisterHandler(HandlerTypeCallbackQueryData, "abc", MatchTypePrefix, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "xabcdef"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "abcdef"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_contains(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id := b.RegisterHandler(HandlerTypeCallbackQueryData, "abc", MatchTypeContains, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "xxabxx"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "xxabcdef"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_regexp(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	re := regexp.MustCompile("^[a-z]+")

	id := b.RegisterHandlerRegexp(HandlerTypeCallbackQueryData, re, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "123abc"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "abcdef"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_invalid_type(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id := b.RegisterHandler(-1, "", -1, nil)

	h := b.handlers[id]

	res := h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "123abc"}})
	if res {
		t.Error("unexpected true result")
	}
}

func TestBot_RegisterUnregisterHandler(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id1 := b.RegisterHandler(HandlerTypeCallbackQueryData, "", MatchTypeExact, nil)
	id2 := b.RegisterHandler(HandlerTypeCallbackQueryData, "", MatchTypeExact, nil)

	if len(b.handlers) != 2 {
		t.Fatalf("unexpected handlers len")
	}
	if _, ok := b.handlers[id1]; !ok {
		t.Fatalf("handler not found")
	}
	if _, ok := b.handlers[id2]; !ok {
		t.Fatalf("handler not found")
	}

	b.UnregisterHandler(id1)
	if len(b.handlers) != 1 {
		t.Fatalf("unexpected handlers len")
	}
	if _, ok := b.handlers[id1]; ok {
		t.Fatalf("handler found")
	}
	if _, ok := b.handlers[id2]; !ok {
		t.Fatalf("handler not found")
	}
}

func Test_match_exact_game(t *testing.T) {
	b := &Bot{
		handlersMx: &sync.RWMutex{},
		handlers:   map[string]handler{},
	}

	id := b.RegisterHandler(HandlerTypeCallbackQueryGameShortName, "xxx", MatchTypeExact, nil)

	h := b.handlers[id]
	u := models.Update{
		ID: 42,
		CallbackQuery: &models.CallbackQuery{
			ID:            "1000",
			GameShortName: "xxx",
		},
	}

	res := h.match(&u)
	if !res {
		t.Error("unexpected true result")
	}
}
