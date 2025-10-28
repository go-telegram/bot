package bot

import (
	"errors"
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
	MatchTypeCommandStartOnly

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

	data, entities, err := getDataFromUpdate(update, h.handlerType)
	if err != nil {
		return false
	}

	switch h.matchType {
	case MatchTypeExact:
		return h.matchExact(data)
	case MatchTypePrefix:
		return h.matchPrefix(data)
	case MatchTypeContains:
		return h.matchContains(data)
	case MatchTypeCommand:
		return h.matchCommand(data, entities)
	case MatchTypeCommandStartOnly:
		return h.matchCommandStartOnly(data, entities)
	case matchTypeRegexp:
		return h.matchRegexp(data)
	default:
		return false
	}
}

func getDataFromUpdate(update *models.Update, handlerType HandlerType) (data string, entities []models.MessageEntity, err error) {
	switch handlerType {
	case HandlerTypeMessageText:
		if update.Message == nil {
			return "", nil, errors.New("message is nil")
		}
		data = update.Message.Text
		entities = update.Message.Entities
	case HandlerTypeCallbackQueryData:
		if update.CallbackQuery == nil {
			return "", nil, errors.New("callback query is nil")
		}
		data = update.CallbackQuery.Data
	case HandlerTypeCallbackQueryGameShortName:
		if update.CallbackQuery == nil {
			return "", nil, errors.New("callback query is nil")
		}
		data = update.CallbackQuery.GameShortName
	case HandlerTypePhotoCaption:
		if update.Message == nil {
			return "", nil, errors.New("message is nil")
		}
		data = update.Message.Caption
		entities = update.Message.CaptionEntities
	}
	return
}

func (h handler) matchExact(data string) bool {
	return data == h.pattern
}

func (h handler) matchPrefix(data string) bool {
	return strings.HasPrefix(data, h.pattern)
}

func (h handler) matchContains(data string) bool {
	return strings.Contains(data, h.pattern)
}

func (h handler) matchRegexp(data string) bool {
	return h.re.Match([]byte(data))
}

func extractCommand(data string, entity models.MessageEntity) string {
	// Checking the correctness of boundaries to avoid panic
	if entity.Offset < 0 || entity.Length <= 1 || entity.Offset+entity.Length > len(data) {
		return ""
	}
	// Skipping the "/" character at the beginning of the command
	return data[entity.Offset+1 : entity.Offset+entity.Length]
}

func (h handler) matchCommand(data string, entities []models.MessageEntity) bool {
	for _, e := range entities {
		if e.Type == models.MessageEntityTypeBotCommand {
			command := extractCommand(data, e)
			if command == h.pattern {
				return true
			}
		}
	}
	return false
}

func (h handler) matchCommandStartOnly(data string, entities []models.MessageEntity) bool {
	for _, e := range entities {
		if e.Type == models.MessageEntityTypeBotCommand && e.Offset == 0 {
			command := extractCommand(data, e)
			if command == h.pattern {
				return true
			}
		}
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
