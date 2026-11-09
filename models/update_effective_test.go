package models

import (
	"encoding/json"
	"testing"
)

func TestUpdateEffective(t *testing.T) {
	// Decode API-shaped payloads to exercise union discriminators and optional
	// fields as callers receive them, including actors distinct from event chats.
	tests := []struct {
		name string
		data string
		user int64
		chat int64
	}{
		{"empty", `{}`, 0, 0},
		{"unknown", `{"future_event":{}}`, 0, 0},
		{"message", `{"message":{"from":{"id":1},"chat":{"id":2},"sender_chat":{"id":3}}}`, 1, 2},
		{"edited", `{"edited_message":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"channel", `{"channel_post":{"chat":{"id":2}}}`, 0, 2},
		{"edited channel", `{"edited_channel_post":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"business connection", `{"business_connection":{"user":{"id":1},"user_chat_id":2}}`, 1, 0},
		{"business message", `{"business_message":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"edited business", `{"edited_business_message":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"deleted business", `{"deleted_business_messages":{"chat":{"id":2}}}`, 0, 2},
		{"reaction", `{"message_reaction":{"user":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"anonymous reaction", `{"message_reaction":{"actor_chat":{"id":3},"chat":{"id":2}}}`, 0, 2},
		{"reaction count", `{"message_reaction_count":{"chat":{"id":2}}}`, 0, 2},
		{"inline query", `{"inline_query":{"from":{"id":1}}}`, 1, 0},
		{"chosen inline", `{"chosen_inline_result":{"from":{"id":1}}}`, 1, 0},
		{"callback", `{"callback_query":{"from":{"id":1},"message":{"date":1,"from":{"id":3},"chat":{"id":2}}}}`, 1, 2},
		{"inaccessible callback", `{"callback_query":{"from":{"id":1},"message":{"date":0,"chat":{"id":2}}}}`, 1, 2},
		{"inline callback", `{"callback_query":{"from":{"id":1},"inline_message_id":"abc"}}`, 1, 0},
		{"widget query", `{"widget_query":{"id":"q","from":{"id":1},"chat":{"id":2},"message_id":5,"chat_instance":"ci","type":{"type":"data"},"data":"d"}}`, 1, 2},
		{"inline widget query", `{"widget_query":{"id":"q","from":{"id":1},"inline_message_id":"abc","chat_instance":"ci","type":{"type":"data"},"data":"d"}}`, 1, 0},
		{"shared widget query", `{"shared_widget_query":{"id":"q","from":{"id":1},"type":{"type":"data"},"data":"d","widget_id":"w"}}`, 1, 0},
		{"shipping", `{"shipping_query":{"from":{"id":1}}}`, 1, 0},
		{"checkout", `{"pre_checkout_query":{"from":{"id":1}}}`, 1, 0},
		{"paid media", `{"purchased_paid_media":{"from":{"id":1}}}`, 1, 0},
		{"poll", `{"poll":{"id":"abc"}}`, 0, 0},
		{"poll answer", `{"poll_answer":{"user":{"id":1}}}`, 1, 0},
		{"anonymous poll answer", `{"poll_answer":{"voter_chat":{"id":3}}}`, 0, 0},
		{"managed bot", `{"managed_bot":{"user":{"id":1},"bot":{"id":3}}}`, 1, 0},
		{"managed bot without user", `{"managed_bot":{"bot":{"id":3}}}`, 0, 0},
		{"guest", `{"guest_message":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"my membership", `{"my_chat_member":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"membership", `{"chat_member":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"join request", `{"chat_join_request":{"from":{"id":1},"chat":{"id":2}}}`, 1, 2},
		{"premium boost", `{"chat_boost":{"chat":{"id":2},"boost":{"source":{"source":"premium","user":{"id":1}}}}}`, 1, 2},
		{"gift boost", `{"chat_boost":{"chat":{"id":2},"boost":{"source":{"source":"gift_code","user":{"id":1}}}}}`, 1, 2},
		{"giveaway boost", `{"chat_boost":{"chat":{"id":2},"boost":{"source":{"source":"giveaway","user":{"id":1}}}}}`, 1, 2},
		{"unclaimed boost", `{"chat_boost":{"chat":{"id":2},"boost":{"source":{"source":"giveaway","is_unclaimed":true}}}}`, 0, 2},
		{"unknown boost", `{"chat_boost":{"chat":{"id":2},"boost":{"source":{"source":"future"}}}}`, 0, 2},
		{"removed boost", `{"removed_chat_boost":{"chat":{"id":2},"source":{"source":"premium","user":{"id":1}}}}`, 1, 2},
		{"subscription", `{"subscription":{"user":{"id":1}}}`, 1, 0},
		{"stopped generation", `{"stopped_message_generation":{"chat":{"id":2}}}`, 0, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Update
			if err := json.Unmarshal([]byte(tt.data), &u); err != nil {
				t.Fatal(err)
			}
			user, chat := u.EffectiveUser(), u.EffectiveChat()
			if tt.user == 0 {
				if user != nil {
					t.Errorf("expected nil user, got %+v", user)
				}
			} else if user == nil || user.ID != tt.user {
				t.Errorf("expected user %d, got %+v", tt.user, user)
			}
			if tt.chat == 0 {
				if chat != nil {
					t.Errorf("expected nil chat, got %+v", chat)
				}
			} else if chat == nil || chat.ID != tt.chat {
				t.Errorf("expected chat %d, got %+v", tt.chat, chat)
			}
		})
	}
}

func TestUpdateEffectiveNilAndIncomplete(t *testing.T) {
	var u *Update
	if u.EffectiveUser() != nil || u.EffectiveChat() != nil {
		t.Fatal("nil update should have no user or chat")
	}
	for _, source := range []ChatBoostSourceType{ChatBoostSourceTypePremium, ChatBoostSourceTypeGiftCode, ChatBoostSourceTypeGiveaway, "future"} {
		u = &Update{ChatBoost: &ChatBoostUpdated{Boost: ChatBoost{Source: ChatBoostSource{Source: source}}}}
		if u.EffectiveUser() != nil {
			t.Errorf("missing boost variant %q should have no user", source)
		}
	}
	for _, kind := range []MaybeInaccessibleMessageType{MaybeInaccessibleMessageTypeMessage, MaybeInaccessibleMessageTypeInaccessibleMessage, 99} {
		u = &Update{CallbackQuery: &CallbackQuery{Message: MaybeInaccessibleMessage{Type: kind}}}
		if u.EffectiveChat() != nil {
			t.Errorf("missing callback message variant %d should have no chat", kind)
		}
	}
}

func TestUpdateEffectiveReturnsOriginalPointers(t *testing.T) {
	u := &Update{Message: &Message{From: &User{ID: 1}, Chat: Chat{ID: 2}}}
	if u.EffectiveUser() != u.Message.From || u.EffectiveChat() != &u.Message.Chat {
		t.Fatal("effective objects must point into the original update")
	}
	u = &Update{CallbackQuery: &CallbackQuery{From: User{ID: 1}, Message: MaybeInaccessibleMessage{Message: &Message{Chat: Chat{ID: 2}}}}}
	if u.EffectiveUser() != &u.CallbackQuery.From || u.EffectiveChat() != &u.CallbackQuery.Message.Message.Chat {
		t.Fatal("callback effective objects must point into the original update")
	}
	u = &Update{WidgetQuery: &WidgetQuery{From: User{ID: 1}, Chat: &Chat{ID: 2}}}
	if u.EffectiveUser() != &u.WidgetQuery.From || u.EffectiveChat() != u.WidgetQuery.Chat {
		t.Fatal("widget query effective objects must point into the original update")
	}
	u = &Update{SharedWidgetQuery: &SharedWidgetQuery{From: User{ID: 1}}}
	if u.EffectiveUser() != &u.SharedWidgetQuery.From || u.EffectiveChat() != nil {
		t.Fatal("shared widget query must return its sender and no chat")
	}
}
