package bot

import (
	"regexp"
	"strings"

	"github.com/go-telegram/bot/models"
)

// HandlerType defines the type of Telegram update content that handler should process.
// Each type corresponds to different data sources within update structure.
type HandlerType int

const (
	HandlerTypeMessageText                HandlerType = iota // Handles text content of regular messages
	HandlerTypeCallbackQueryData                             // Handles callback data from inline keyboard buttons
	HandlerTypeCallbackQueryGameShortName                    // Handles game short names from game callback queries
	HandlerTypePhotoCaption                                  // Handles caption text of photo messages
)

// MatchType defines how the handler pattern should be matched against incoming content.
// Different match types provide various levels of flexibility for content filtering.
type MatchType int

const (
	MatchTypeExact            MatchType = iota // Exact string match (content == pattern)
	MatchTypePrefix                            // Prefix match (content starts with pattern)
	MatchTypeContains                          // Contains match (pattern found anywhere in content)
	MatchTypeCommand                           // Bot command match (finds /pattern as bot command)
	MatchTypeCommandStartOnly                  // Bot command at message start only (strict command matching)

	// Internal match types (not for public use)
	matchTypeRegexp // Regular expression matching (use RegisterHandlerRegexp)
	matchTypeFunc   // Custom function matching (use RegisterHandlerMatchFunc)
)

// handler represents internal handler configuration with matching logic and execution function.
// Contains all necessary data for determining if update should be processed by this handler.
type handler struct {
	id          string      // Unique handler identifier for registration/unregistration
	handlerType HandlerType // Type of content to extract from update
	matchType   MatchType   // How to match extracted content against pattern
	handler     HandlerFunc // Function to execute when match is found

	// Pattern matching fields (used based on matchType)
	pattern   string         // String pattern for basic matching
	re        *regexp.Regexp // Compiled regex for regexp matching
	matchFunc MatchFunc      // Custom function for function-based matching
}

func (h handler) match(update *models.Update) bool {
	if h.matchType == matchTypeFunc {
		return h.matchFunc(update)
	}

	var data string
	var entities []models.MessageEntity

	switch h.handlerType {
	case HandlerTypeMessageText:
		if update.Message == nil {
			return false
		}
		data = update.Message.Text
		entities = update.Message.Entities
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
		entities = update.Message.CaptionEntities
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
	if h.matchType == MatchTypeCommand {
		for _, e := range entities {
			if e.Type == models.MessageEntityTypeBotCommand {
				if data[e.Offset+1:e.Offset+e.Length] == h.pattern {
					return true
				}
			}
		}
	}
	if h.matchType == MatchTypeCommandStartOnly {
		for _, e := range entities {
			if e.Type == models.MessageEntityTypeBotCommand {
				if e.Offset == 0 && data[e.Offset+1:e.Offset+e.Length] == h.pattern {
					return true
				}
			}
		}
	}
	if h.matchType == matchTypeRegexp {
		return h.re.Match([]byte(data))
	}
	return false
}

// RegisterHandlerMatchFunc registers a handler with custom matching function.
// This provides maximum flexibility for complex matching logic that cannot be expressed with standard match types.
// Returns handler ID that can be used to unregister the handler later.
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

// RegisterHandlerRegexp registers a handler with regular expression pattern matching.
// The regex will be applied to content extracted based on handlerType.
// Returns handler ID that can be used to unregister the handler later.
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

// RegisterHandler registers a handler with specified pattern and match type.
// This is the most commonly used method for registering handlers with standard matching logic.
// Returns handler ID that can be used to unregister the handler later.
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

// UnregisterHandler removes a previously registered handler by its ID.
// The ID is returned when registering handlers with any of the Register* methods.
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
