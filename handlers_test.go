package bot

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/go-telegram/bot/models"
)

func findHandler(b *Bot, id string) *handler {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	for _, h := range b.handlers {
		if h.id == id {
			return &h
		}
	}

	return nil
}

func Test_match_func(t *testing.T) {
	b := &Bot{}

	var called bool

	id := b.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		called = true
		if update.ID != 42 {
			t.Error("invalid update id")
		}
		return true
	}, nil)

	h := findHandler(b, id)

	res := h.match(&models.Update{ID: 42})
	if !called {
		t.Error("not called")
	}
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_exact(t *testing.T) {
	b := &Bot{}

	id := b.RegisterHandler(HandlerTypeMessageText, "xxx", MatchTypeExact, nil)

	h := findHandler(b, id)

	res := h.match(&models.Update{Message: &models.Message{Text: "zzz"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{Message: &models.Message{Text: "xxx"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_caption_exact(t *testing.T) {
	b := &Bot{}

	id := b.RegisterHandler(HandlerTypePhotoCaption, "xxx", MatchTypeExact, nil)

	h := findHandler(b, id)

	res := h.match(&models.Update{Message: &models.Message{Caption: "zzz"}})
	if res {
		t.Error("unexpected true result")
	}

	res = h.match(&models.Update{Message: &models.Message{Caption: "xxx"}})
	if !res {
		t.Error("unexpected false result")
	}
}

func Test_match_prefix(t *testing.T) {
	b := &Bot{}

	id := b.RegisterHandler(HandlerTypeCallbackQueryData, "abc", MatchTypePrefix, nil)

	h := findHandler(b, id)

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
	b := &Bot{}

	id := b.RegisterHandler(HandlerTypeCallbackQueryData, "abc", MatchTypeContains, nil)

	h := findHandler(b, id)

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
	b := &Bot{}

	re := regexp.MustCompile("^[a-z]+")

	id := b.RegisterHandlerRegexp(HandlerTypeCallbackQueryData, re, nil)

	h := findHandler(b, id)

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
	b := &Bot{}

	id := b.RegisterHandler(-1, "", -1, nil)

	h := findHandler(b, id)

	res := h.match(&models.Update{CallbackQuery: &models.CallbackQuery{Data: "123abc"}})
	if res {
		t.Error("unexpected true result")
	}
}

func TestBot_RegisterUnregisterHandler(t *testing.T) {
	b := &Bot{}

	id1 := b.RegisterHandler(HandlerTypeCallbackQueryData, "", MatchTypeExact, nil)
	id2 := b.RegisterHandler(HandlerTypeCallbackQueryData, "", MatchTypeExact, nil)

	if len(b.handlers) != 2 {
		t.Fatalf("unexpected handlers len")
	}
	if h := findHandler(b, id1); h == nil {
		t.Fatalf("handler not found")
	}
	if h := findHandler(b, id2); h == nil {
		t.Fatalf("handler not found")
	}

	b.UnregisterHandler(id1)
	if len(b.handlers) != 1 {
		t.Fatalf("unexpected handlers len")
	}
	if h := findHandler(b, id1); h != nil {
		t.Fatalf("handler found")
	}
	if h := findHandler(b, id2); h == nil {
		t.Fatalf("handler not found")
	}
}

func Test_match_exact_game(t *testing.T) {
	b := &Bot{}

	id := b.RegisterHandler(HandlerTypeCallbackQueryGameShortName, "xxx", MatchTypeExact, nil)

	h := findHandler(b, id)
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

func Test_match_command_start(t *testing.T) {
	t.Run("anywhere 1, yes", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommand, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "/foo",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if !res {
			t.Error("unexpected result")
		}
	})

	t.Run("anywhere 2, yes", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommand, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "a /foo",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 2, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if !res {
			t.Error("unexpected result")
		}
	})

	t.Run("anywhere 3, no", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommand, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "a /bar",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 2, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if res {
			t.Error("unexpected result")
		}
	})

	t.Run("start 1, yes", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommandStartOnly, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "/foo",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if !res {
			t.Error("unexpected result")
		}
	})

	t.Run("start 2, no", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommandStartOnly, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "a /foo",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 2, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if res {
			t.Error("unexpected result")
		}
	})

	t.Run("start 3, no", func(t *testing.T) {
		b := &Bot{}

		id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommandStartOnly, nil)

		h := findHandler(b, id)
		u := models.Update{
			ID: 42,
			Message: &models.Message{
				Text: "/bar",
				Entities: []models.MessageEntity{
					{Type: models.MessageEntityTypeBotCommand, Offset: 2, Length: 4},
				},
			},
		}

		res := h.match(&u)
		if res {
			t.Error("unexpected result")
		}
	})
}

func Test_match_NilUpdateMessageIsFalse(t *testing.T) {
	b := &Bot{}
	id := b.RegisterHandler(HandlerTypeMessageText, "foo", MatchTypeCommand, nil)
	h := findHandler(b, id)

	u := models.Update{
		ID:      42,
		Message: nil,
	}

	res := h.match(&u)
	if res {
		t.Error("want 'false', but got 'true'")
	}
}

func Test_getDataFromUpdate(t *testing.T) {
	tests := []struct {
		name         string
		update       *models.Update
		handlerType  HandlerType
		wantData     string
		wantEntities []models.MessageEntity
		wantOK       bool
	}{
		{
			name: "HandlerTypeMessageText - valid message, is ok",
			update: &models.Update{
				Message: &models.Message{
					Text: "Hello, world!",
					Entities: []models.MessageEntity{
						{Type: models.MessageEntityTypeBold, Offset: 0, Length: 5},
					},
				},
			},
			handlerType: HandlerTypeMessageText,
			wantData:    "Hello, world!",
			wantEntities: []models.MessageEntity{
				{Type: models.MessageEntityTypeBold, Offset: 0, Length: 5},
			},
			wantOK: true,
		},
		{
			name: "HandlerTypeMessageText - nil message, is unavailable",
			update: &models.Update{
				Message: nil,
			},
			handlerType:  HandlerTypeMessageText,
			wantData:     "",
			wantEntities: nil,
			wantOK:       false,
		},
		{
			name: "HandlerTypeMessageText - empty message, is ok",
			update: &models.Update{
				Message: &models.Message{
					Text:     "",
					Entities: nil,
				},
			},
			handlerType:  HandlerTypeMessageText,
			wantData:     "",
			wantEntities: nil,
			wantOK:       true,
		},
		{
			name: "HandlerTypeCallbackQueryData - valid callback query, is ok",
			update: &models.Update{
				CallbackQuery: &models.CallbackQuery{
					Data: "callback_data",
				},
			},
			handlerType:  HandlerTypeCallbackQueryData,
			wantData:     "callback_data",
			wantEntities: nil,
			wantOK:       true,
		},
		{
			name: "HandlerTypeCallbackQueryData - nil callback query, is unavailable",
			update: &models.Update{
				CallbackQuery: nil,
			},
			handlerType:  HandlerTypeCallbackQueryData,
			wantData:     "",
			wantEntities: nil,
			wantOK:       false,
		},
		{
			name: "HandlerTypeCallbackQueryData - empty data, is ok",
			update: &models.Update{
				CallbackQuery: &models.CallbackQuery{
					Data: "",
				},
			},
			handlerType:  HandlerTypeCallbackQueryData,
			wantData:     "",
			wantEntities: nil,
			wantOK:       true,
		},
		{
			name: "HandlerTypeCallbackQueryGameShortName - valid game short name, is ok",
			update: &models.Update{
				CallbackQuery: &models.CallbackQuery{
					GameShortName: "snake_game",
				},
			},
			handlerType:  HandlerTypeCallbackQueryGameShortName,
			wantData:     "snake_game",
			wantEntities: nil,
			wantOK:       true,
		},
		{
			name: "HandlerTypeCallbackQueryGameShortName - nil callback query, is unavailable",
			update: &models.Update{
				CallbackQuery: nil,
			},
			handlerType:  HandlerTypeCallbackQueryGameShortName,
			wantData:     "",
			wantEntities: nil,
			wantOK:       false,
		},
		{
			name: "HandlerTypePhotoCaption - valid photo caption, is ok",
			update: &models.Update{
				Message: &models.Message{
					Caption: "Photo caption",
					CaptionEntities: []models.MessageEntity{
						{Type: models.MessageEntityTypeItalic, Offset: 6, Length: 7},
					},
				},
			},
			handlerType: HandlerTypePhotoCaption,
			wantData:    "Photo caption",
			wantEntities: []models.MessageEntity{
				{Type: models.MessageEntityTypeItalic, Offset: 6, Length: 7},
			},
			wantOK: true,
		},
		{
			name: "HandlerTypePhotoCaption - nil message, is unavailable",
			update: &models.Update{
				Message: nil,
			},
			handlerType:  HandlerTypePhotoCaption,
			wantData:     "",
			wantEntities: nil,
			wantOK:       false,
		},
		{
			name: "HandlerTypePhotoCaption - empty caption, is ok",
			update: &models.Update{
				Message: &models.Message{
					Caption:         "",
					CaptionEntities: nil,
				},
			},
			handlerType:  HandlerTypePhotoCaption,
			wantData:     "",
			wantEntities: nil,
			wantOK:       true,
		},
		{
			name: "Invalid handler type returns empty values, is ok",
			update: &models.Update{
				Message: &models.Message{
					Text: "some text",
				},
			},
			handlerType:  HandlerType(999), // Unknown type
			wantData:     "",
			wantEntities: nil,
			wantOK:       true, // Func just returns empty values for unknown types
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, entities, ok := getDataFromUpdate(tt.update, tt.handlerType)

			if ok != tt.wantOK {
				t.Errorf("getDataFromUpdate() ok = %v, want %v", ok, tt.wantOK)
			}

			if data != tt.wantData {
				t.Errorf("want data %q, got %q", tt.wantData, data)
			}

			if !reflect.DeepEqual(entities, tt.wantEntities) {
				t.Errorf("entities mismatch:\nwant: %+v\ngot:  %+v", tt.wantEntities, entities)
			}
		})
	}
}

func TestHandler_matchExact(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		data      string
		wantMatch bool
	}{
		{
			name:      "exact match",
			pattern:   "test",
			data:      "test",
			wantMatch: true,
		},
		{
			name:      "no match",
			pattern:   "test",
			data:      "testing",
			wantMatch: false,
		},
		{
			name:      "empty pattern",
			pattern:   "",
			data:      "",
			wantMatch: true,
		},
		{
			name:      "empty data",
			pattern:   "test",
			data:      "",
			wantMatch: false,
		},
		{
			name:      "case sensitive",
			pattern:   "Test",
			data:      "test",
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{pattern: tt.pattern}
			result := h.matchExact(tt.data)
			if result != tt.wantMatch {
				t.Errorf("matchExact() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func TestHandler_matchPrefix(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		data      string
		wantMatch bool
	}{
		{
			name:      "has prefix",
			pattern:   "test",
			data:      "testing",
			wantMatch: true,
		},
		{
			name:      "exact match",
			pattern:   "test",
			data:      "test",
			wantMatch: true,
		},
		{
			name:      "no prefix",
			pattern:   "test",
			data:      "notest",
			wantMatch: false,
		},
		{
			name:      "empty pattern",
			pattern:   "",
			data:      "anything",
			wantMatch: true,
		},
		{
			name:      "empty data",
			pattern:   "test",
			data:      "",
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{pattern: tt.pattern}
			result := h.matchPrefix(tt.data)
			if result != tt.wantMatch {
				t.Errorf("matchPrefix() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func TestHandler_matchContains(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		data      string
		wantMatch bool
	}{
		{
			name:      "contains",
			pattern:   "test",
			data:      "atestb",
			wantMatch: true,
		},
		{
			name:      "exact match",
			pattern:   "test",
			data:      "test",
			wantMatch: true,
		},
		{
			name:      "no contains",
			pattern:   "test",
			data:      "nothing",
			wantMatch: false,
		},
		{
			name:      "empty pattern",
			pattern:   "",
			data:      "anything",
			wantMatch: true,
		},
		{
			name:      "empty data",
			pattern:   "test",
			data:      "",
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{pattern: tt.pattern}
			result := h.matchContains(tt.data)
			if result != tt.wantMatch {
				t.Errorf("matchContains() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func TestHandler_matchRegexp(t *testing.T) {
	tests := []struct {
		name      string
		regexp    string
		data      string
		wantMatch bool
	}{
		{
			name:      "simple match",
			regexp:    "^test$",
			data:      "test",
			wantMatch: true,
		},
		{
			name:      "no match",
			regexp:    "^test$",
			data:      "testing",
			wantMatch: false,
		},
		{
			name:      "partial match",
			regexp:    "test",
			data:      "testing",
			wantMatch: true,
		},
		{
			name:      "digit pattern",
			regexp:    "\\d+",
			data:      "123abc",
			wantMatch: true,
		},
		{
			name:      "no digits",
			regexp:    "\\d+",
			data:      "abc",
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := regexp.MustCompile(tt.regexp)
			h := handler{re: re}
			result := h.matchRegexp(tt.data)
			if result != tt.wantMatch {
				t.Errorf("matchRegexp() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func TestExtractCommand(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		entity models.MessageEntity
		want   string
		wantOK bool
	}{
		{
			name: "valid command",
			data: "/start arg1",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 0,
				Length: 6,
			},
			want:   "start",
			wantOK: true,
		},
		{
			name: "command in middle",
			data: "text /help more",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 5,
				Length: 5,
			},
			want:   "help",
			wantOK: true,
		},
		{
			name: "invalid offset negative",
			data: "/start",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: -1,
				Length: 6,
			},
			want: "",
		},
		{
			name: "invalid length too long",
			data: "/start",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 0,
				Length: 100,
			},
			want: "",
		},
		{
			name: "length too small",
			data: "/",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 0,
				Length: 1,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := extractCommand(tt.data, tt.entity)
			if ok != tt.wantOK {
				t.Errorf("extractCommand() ok = %v, want %v", ok, tt.wantOK)
			}
			if result != tt.want {
				t.Errorf("extractCommand() = %q, want %q", result, tt.want)
			}
		})
	}
}

func TestHandler_matchCommand(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		data      string
		entities  []models.MessageEntity
		wantMatch bool
	}{
		{
			name:    "command match anywhere",
			pattern: "start",
			data:    "text /start arg1",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 5,
					Length: 6,
				},
			},
			wantMatch: true,
		},
		{
			name:    "no command match",
			pattern: "help",
			data:    "text /start arg1",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 5,
					Length: 6,
				},
			},
			wantMatch: false,
		},
		{
			name:      "no entities",
			pattern:   "start",
			data:      "/start",
			entities:  []models.MessageEntity{},
			wantMatch: false,
		},
		{
			name:    "multiple commands, one matches",
			pattern: "help",
			data:    "/start /help /stop",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 0,
					Length: 6,
				},
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 7,
					Length: 5,
				},
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 13,
					Length: 5,
				},
			},
			wantMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{pattern: tt.pattern}
			result := h.matchCommand(tt.data, tt.entities)
			if result != tt.wantMatch {
				t.Errorf("matchCommand() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func TestHandler_matchCommandStartOnly(t *testing.T) {
	tests := []struct {
		name      string
		pattern   string
		data      string
		entities  []models.MessageEntity
		wantMatch bool
	}{
		{
			name:    "command at start",
			pattern: "start",
			data:    "/start arg1",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 0,
					Length: 6,
				},
			},
			wantMatch: true,
		},
		{
			name:    "command not at start",
			pattern: "start",
			data:    "text /start arg1",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 5,
					Length: 6,
				},
			},
			wantMatch: false,
		},
		{
			name:    "wrong command at start",
			pattern: "help",
			data:    "/start arg1",
			entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBotCommand,
					Offset: 0,
					Length: 6,
				},
			},
			wantMatch: false,
		},
		{
			name:      "no entities",
			pattern:   "start",
			data:      "/start",
			entities:  []models.MessageEntity{},
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{pattern: tt.pattern}
			result := h.matchCommandStartOnly(tt.data, tt.entities)
			if result != tt.wantMatch {
				t.Errorf("matchCommandStartOnly() = %v, want %v", result, tt.wantMatch)
			}
		})
	}
}

func Test_match_invalidCommandEntities(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	tests := []struct {
		name   string
		offset int
		length int
	}{
		{name: "negative offset", offset: -1, length: 6},
		{name: "offset past end", offset: 7, length: 2},
		{name: "offset at end", offset: 6, length: 2},
		{name: "negative length", length: -1},
		{name: "zero length", length: 0},
		{name: "slash only", length: 1},
		{name: "length past end", length: 100},
		{name: "offset addition overflow", offset: maxInt, length: 2},
		{name: "length addition overflow", offset: 1, length: maxInt},
	}

	for _, handlerType := range []HandlerType{HandlerTypeMessageText, HandlerTypePhotoCaption} {
		for _, matchType := range []MatchType{MatchTypeCommand, MatchTypeCommandStartOnly} {
			for _, pattern := range []string{"", "start"} {
				for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
						entities := []models.MessageEntity{{
							Type:   models.MessageEntityTypeBotCommand,
							Offset: tt.offset,
							Length: tt.length,
						}}
						u := &models.Update{Message: &models.Message{
							Text:            "/start",
							Entities:        entities,
							Caption:         "/start",
							CaptionEntities: entities,
						}}
						h := handler{handlerType: handlerType, matchType: matchType, pattern: pattern}
						if h.match(u) {
							t.Errorf("invalid command matched: handlerType=%d matchType=%d pattern=%q", handlerType, matchType, pattern)
						}
					})
				}
			}
		}
	}
}

func Test_match_commandSkipsInvalidEntities(t *testing.T) {
	entities := []models.MessageEntity{
		{Type: models.MessageEntityTypeBold, Offset: 0, Length: 6},
		{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: 100},
		{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: 6},
	}
	u := &models.Update{Message: &models.Message{
		Text:            "/start",
		Entities:        entities,
		Caption:         "/start",
		CaptionEntities: entities,
	}}
	for _, handlerType := range []HandlerType{HandlerTypeMessageText, HandlerTypePhotoCaption} {
		for _, matchType := range []MatchType{MatchTypeCommand, MatchTypeCommandStartOnly} {
			h := handler{handlerType: handlerType, matchType: matchType, pattern: "start"}
			if !h.match(u) {
				t.Errorf("valid command after invalid entity did not match: handlerType=%d matchType=%d", handlerType, matchType)
			}
		}
	}
}

func Test_match_missingUpdateData(t *testing.T) {
	for _, handlerType := range []HandlerType{HandlerTypeMessageText, HandlerTypePhotoCaption, HandlerTypeCallbackQueryData, HandlerTypeCallbackQueryGameShortName} {
		for _, matchType := range []MatchType{MatchTypeExact, MatchTypePrefix, MatchTypeContains, matchTypeRegexp} {
			h := handler{handlerType: handlerType, matchType: matchType, re: regexp.MustCompile("")}
			if h.match(&models.Update{}) {
				t.Errorf("missing data matched an empty pattern: handlerType=%d matchType=%d", handlerType, matchType)
			}
		}
	}
}

func Test_match_unknownHandlerTypePreservesEmptyData(t *testing.T) {
	for _, matchType := range []MatchType{MatchTypeExact, MatchTypePrefix, MatchTypeContains, matchTypeRegexp} {
		h := handler{handlerType: HandlerType(999), matchType: matchType, re: regexp.MustCompile("")}
		if !h.match(&models.Update{}) {
			t.Errorf("unknown handler type no longer matches an empty pattern: matchType=%d", matchType)
		}
	}
}
