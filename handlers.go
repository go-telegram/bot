package bot

import (
	"regexp"
	"strings"

	"github.com/go-telegram/bot/models"
)

// HandlerType selects the update data that a handler matches against.
type HandlerType int

const (
	HandlerTypeMessageText                HandlerType = iota // Text of Update.Message
	HandlerTypeCallbackQueryData                             // Data of Update.CallbackQuery
	HandlerTypeCallbackQueryGameShortName                    // GameShortName of Update.CallbackQuery
	HandlerTypePhotoCaption                                  // Caption of Update.Message, for any message type
	HandlerTypeInlineQuery                                   // Query of Update.InlineQuery
	HandlerTypeWidgetQueryData                               // Data of Update.WidgetQuery
	HandlerTypeSharedWidgetQueryData                         // Data of Update.SharedWidgetQuery
)

// MatchType defines how a handler pattern is compared with the update data.
type MatchType int

const (
	MatchTypeExact    MatchType = iota // Data equals pattern
	MatchTypePrefix                    // Data starts with pattern
	MatchTypeContains                  // Data contains pattern
	// MatchTypeCommand matches a bot command anywhere in a message text or
	// caption, so it works only with HandlerTypeMessageText and
	// HandlerTypePhotoCaption. Specify the command without the slash, such as
	// "start". It also matches /start@bot_username once GetMe has succeeded.
	MatchTypeCommand
	// MatchTypeCommandStartOnly is like MatchTypeCommand, but the command must
	// be at the start of the text or caption.
	MatchTypeCommandStartOnly

	// Internal match types (not for public use)
	matchTypeRegexp // Regular expression matching (use RegisterHandlerRegexp)
	matchTypeFunc   // Custom function matching (use RegisterHandlerMatchFunc)
)

// handler is a registered update handler with its matching rules.
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

// match reports whether the handler accepts update. botUsername is the bot's
// username, or empty if unknown.
func (h handler) match(update *models.Update, botUsername string) bool {
	if h.matchType == matchTypeFunc {
		return h.matchFunc(update)
	}

	data, entities, ok := getDataFromUpdate(update, h.handlerType)
	if !ok {
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
		return h.matchCommand(data, entities, botUsername)
	case MatchTypeCommandStartOnly:
		return h.matchCommandStartOnly(data, entities, botUsername)
	case matchTypeRegexp:
		return h.matchRegexp(data)
	default:
		return false
	}
}

func getDataFromUpdate(update *models.Update, handlerType HandlerType) (data string, entities []models.MessageEntity, ok bool) {
	switch handlerType {
	case HandlerTypeMessageText:
		if update.Message == nil {
			return "", nil, false
		}
		data = update.Message.Text
		entities = update.Message.Entities
	case HandlerTypeCallbackQueryData:
		if update.CallbackQuery == nil {
			return "", nil, false
		}
		data = update.CallbackQuery.Data
	case HandlerTypeCallbackQueryGameShortName:
		if update.CallbackQuery == nil {
			return "", nil, false
		}
		data = update.CallbackQuery.GameShortName
	case HandlerTypePhotoCaption:
		if update.Message == nil {
			return "", nil, false
		}
		data = update.Message.Caption
		entities = update.Message.CaptionEntities
	case HandlerTypeInlineQuery:
		if update.InlineQuery == nil {
			return "", nil, false
		}
		data = update.InlineQuery.Query
	case HandlerTypeWidgetQueryData:
		if update.WidgetQuery == nil {
			return "", nil, false
		}
		data = update.WidgetQuery.Data
	case HandlerTypeSharedWidgetQueryData:
		if update.SharedWidgetQuery == nil {
			return "", nil, false
		}
		data = update.SharedWidgetQuery.Data
	}
	return data, entities, true
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

func extractCommand(data string, entity models.MessageEntity) (string, bool) {
	// Check the offset before subtracting, avoiding overflow in offset + length.
	// A string never has more UTF-16 code units than bytes, so an entity that
	// fails this check cannot fit in the text either.
	if entity.Offset < 0 || entity.Offset > len(data) || entity.Length <= 1 || entity.Length > len(data)-entity.Offset {
		return "", false
	}
	// Telegram counts entity offsets and lengths in UTF-16 code units.
	from, okFrom := utf16Index(data, entity.Offset)
	to, okTo := utf16Index(data[from:], entity.Length)
	if !okFrom || !okTo {
		return "", false
	}
	// Skip the "/" character at the beginning of the command.
	return data[from+1 : from+to], true
}

// utf16Index returns the byte index of the n-th UTF-16 code unit in s. It
// reports false when n is negative, falls inside a surrogate pair or is past
// the end of s. An invalid UTF-8 byte counts as one unit, as U+FFFD would.
func utf16Index(s string, n int) (int, bool) {
	units := 0
	for i, r := range s {
		if units >= n {
			return i, units == n
		}
		units++
		if r > 0xFFFF {
			// Outside the BMP, a rune takes a surrogate pair.
			units++
		}
	}
	return len(s), units == n
}

func (h handler) matchCommand(data string, entities []models.MessageEntity, botUsername string) bool {
	for _, e := range entities {
		if e.Type == models.MessageEntityTypeBotCommand {
			command, ok := extractCommand(data, e)
			if ok && h.isCommand(command, botUsername) {
				return true
			}
		}
	}
	return false
}

func (h handler) matchCommandStartOnly(data string, entities []models.MessageEntity, botUsername string) bool {
	for _, e := range entities {
		if e.Type == models.MessageEntityTypeBotCommand && e.Offset == 0 {
			command, ok := extractCommand(data, e)
			if ok && h.isCommand(command, botUsername) {
				return true
			}
		}
	}
	return false
}

// isCommand reports whether command, without its leading "/", is the handler's
// command. In groups, clients address a command to a bot by appending its
// username, as in /start@bot_username. Telegram usernames are case-insensitive.
func (h handler) isCommand(command, botUsername string) bool {
	if command == h.pattern {
		return true
	}
	if h.pattern == "" || botUsername == "" {
		return false
	}
	name, target, addressed := strings.Cut(command, "@")
	return addressed && name == h.pattern && strings.EqualFold(target, botUsername)
}

// RegisterHandlerMatchFunc registers f for updates accepted by matchFunc.
// Handlers are checked in registration order, and only the first match runs.
// Middlewares in m wrap f. Returns an ID that can be passed to UnregisterHandler.
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

// RegisterHandlerRegexp registers f for updates whose data, selected by
// handlerType, matches re. Handlers are checked in registration order, and only
// the first match runs. Middlewares in m wrap f. Returns an ID that can be
// passed to UnregisterHandler.
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

// RegisterHandlerInlineQuery registers a handler for all inline queries, including
// empty queries. For matching query text, use RegisterHandler or RegisterHandlerRegexp
// with HandlerTypeInlineQuery. Returns an ID that can be passed to UnregisterHandler.
func (b *Bot) RegisterHandlerInlineQuery(f HandlerFunc, m ...Middleware) string {
	return b.RegisterHandler(HandlerTypeInlineQuery, "", MatchTypePrefix, f, m...)
}

// RegisterHandler registers f for updates whose data, selected by handlerType,
// matches pattern according to matchType. Handlers are checked in registration
// order, and only the first match runs. Middlewares in m wrap f. Returns an ID
// that can be passed to UnregisterHandler.
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

// UnregisterHandler removes the handler with the given ID. An unknown ID is ignored.
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
