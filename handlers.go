package bot

import (
	"regexp"
	"strings"
)

type HandlerType int

const (
	HandlerTypeMessageText HandlerType = iota
	HandlerTypeCallbackQueryData
)

type MatchType int

const (
	MatchTypeExact MatchType = iota
	MatchTypePrefix
	MatchTypeContains
	MatchTypeRegexp
)

type handler struct {
	handlerType HandlerType
	matchType   MatchType
	pattern     string
	handler     HandlerFunc
	re          *regexp.Regexp
}

func (h handler) match(data string) bool {
	if h.matchType == MatchTypeExact {
		return data == h.pattern
	}
	if h.matchType == MatchTypePrefix {
		return strings.HasPrefix(data, h.pattern)
	}
	if h.matchType == MatchTypeContains {
		return strings.Contains(data, h.pattern)
	}
	if h.matchType == MatchTypeRegexp {
		return h.re.Match([]byte(data))
	}
	return false
}

func (b *Bot) RegisterHandler(handlerType HandlerType, pattern string, matchType MatchType, f HandlerFunc) (string, error) {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	id := RandomString(16)

	h := handler{
		handlerType: handlerType,
		matchType:   matchType,
		pattern:     pattern,
		handler:     f,
	}

	var err error

	if matchType == MatchTypeRegexp {
		h.re, err = regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
	}

	b.handlers[id] = h

	return id, nil
}

func (b *Bot) UnregisterHandler(id string) {
	b.handlersMx.Lock()
	defer b.handlersMx.Unlock()

	delete(b.handlers, id)
}
