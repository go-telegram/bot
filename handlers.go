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
	HandlerTypePhotoCaption
)

type MatchType int

const (
	MatchTypeExact MatchType = iota
	MatchTypePrefix
	MatchTypeContains
	MatchTypeCommand

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

var CommandRegexp = regexp.MustCompile(`(?s)^/(\w+)(?:@(\w+))?(?:\s+(.+?)\s*)?$`)

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
		if update.CallbackQuery == nil {
			return false
		}
		data = update.CallbackQuery.GameShortName
	case HandlerTypePhotoCaption:
		if update.Message == nil {
			return false
		}
		data = update.Message.Caption
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
	if h.matchType == MatchTypeCommand {
		matches := CommandRegexp.FindStringSubmatch(data)
		if len(matches) != 4 {
			return false
		}

		return strings.EqualFold(matches[1], h.pattern)
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
