package bot

import (
	"regexp"
	"strings"

	"github.com/go-telegram/bot/models"
)

type HandlerType int

const (
	HandlerTypeMessageText HandlerType = iota
	HandlerTypeCallbackQueryData
	HandlerTypeCallbackQueryGameShortName
)

type MatchType int

const (
	MatchTypeExact MatchType = iota
	MatchTypePrefix
	MatchTypeContains

	matchTypeRegexp
	matchTypeFunc
)

type handler struct {
	id          string
	handlerType HandlerType
	matchType   MatchType
	handler     HandlerFunc

	pattern   string
	re        *regexp.Regexp
	matchFunc MatchFunc
}

func (h handler) match(update *models.Update) bool {
	if h.matchType == matchTypeFunc {
		return h.matchFunc(update)
	}

	var data string

	switch h.handlerType {
	case HandlerTypeMessageText:
		if update.Message == nil {
			return false
		}
		data = update.Message.Text
	case HandlerTypeCallbackQueryData:
		if update.CallbackQuery == nil {
			return false
		}
		data = update.CallbackQuery.Data
	case HandlerTypeCallbackQueryGameShortName:
		data = update.CallbackQuery.GameShortName
	}

	if h.matchType == MatchTypeExact {
		return data == h.pattern
	}
	if h.matchType == MatchTypePrefix {
		return strings.HasPrefix(data, h.pattern)
	}
	if h.matchType == MatchTypeContains {
		return strings.Contains(data, h.pattern)
	}
	if h.matchType == matchTypeRegexp {
		return h.re.Match([]byte(data))
	}
	return false
}

func (b *Bot) RegisterHandlerMatchFunc(matchFunc MatchFunc, f HandlerFunc, m ...Middleware) string {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	id := RandomString(16)

	h := handler{
		id:        id,
		matchType: matchTypeFunc,
		matchFunc: matchFunc,
		handler:   applyMiddlewares(f, m...),
	}

	b.handlers = append(b.handlers, h)

	return id
}

func (b *Bot) RegisterHandlerRegexp(handlerType HandlerType, re *regexp.Regexp, f HandlerFunc, m ...Middleware) string {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	id := RandomString(16)

	h := handler{
		id:          id,
		handlerType: handlerType,
		matchType:   matchTypeRegexp,
		re:          re,
		handler:     applyMiddlewares(f, m...),
	}

	b.handlers = append(b.handlers, h)

	return id
}

func (b *Bot) RegisterHandler(handlerType HandlerType, pattern string, matchType MatchType, f HandlerFunc, m ...Middleware) string {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	id := RandomString(16)

	h := handler{
		id:          id,
		handlerType: handlerType,
		matchType:   matchType,
		pattern:     pattern,
		handler:     applyMiddlewares(f, m...),
	}

	b.handlers = append(b.handlers, h)

	return id
}

func (b *Bot) UnregisterHandler(id string) {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	for i, h := range b.handlers {
		if h.id == id {
			b.handlers = append(b.handlers[:i], b.handlers[i+1:]...)
			return
		}
	}
}
