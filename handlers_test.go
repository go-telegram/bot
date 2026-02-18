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
		wantError    bool
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
			wantError: false,
		},
		{
			name: "HandlerTypeMessageText - nil message, is error",
			update: &models.Update{
				Message: nil,
			},
			handlerType:  HandlerTypeMessageText,
			wantData:     "",
			wantEntities: nil,
			wantError:    true,
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
			wantError:    false,
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
			wantError:    false,
		},
		{
			name: "HandlerTypeCallbackQueryData - nil callback query, is error",
			update: &models.Update{
				CallbackQuery: nil,
			},
			handlerType:  HandlerTypeCallbackQueryData,
			wantData:     "",
			wantEntities: nil,
			wantError:    true,
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
			wantError:    false,
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
			wantError:    false,
		},
		{
			name: "HandlerTypeCallbackQueryGameShortName - nil callback query, is error",
			update: &models.Update{
				CallbackQuery: nil,
			},
			handlerType:  HandlerTypeCallbackQueryGameShortName,
			wantData:     "",
			wantEntities: nil,
			wantError:    true,
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
			wantError: false,
		},
		{
			name: "HandlerTypePhotoCaption - nil message, is error",
			update: &models.Update{
				Message: nil,
			},
			handlerType:  HandlerTypePhotoCaption,
			wantData:     "",
			wantEntities: nil,
			wantError:    true,
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
			wantError:    false,
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
			wantError:    false, // Func just returns empty values for unknown types
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, entities, err := getDataFromUpdate(tt.update, tt.handlerType)

			if tt.wantError && err == nil {
				t.Errorf("want error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("unexpected error: %v", err)
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
	}{
		{
			name: "valid command",
			data: "/start arg1",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 0,
				Length: 6,
			},
			want: "start",
		},
		{
			name: "command in middle",
			data: "text /help more",
			entity: models.MessageEntity{
				Type:   models.MessageEntityTypeBotCommand,
				Offset: 5,
				Length: 5,
			},
			want: "help",
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
			result := extractCommand(tt.data, tt.entity)
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
