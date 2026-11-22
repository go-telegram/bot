package models

import (
	"encoding/json"
	"testing"
)

func TestBotFullInfo(t *testing.T) {
	const src = `{"id":42,"is_bot":true,"first_name":"Bot","username":"foo","active_usernames":["foo","foo_bot"],"supports_inline_queries":true,"inline_query_placeholder":"search","needs_inline_query_location":true,"has_main_web_app":true,"has_topic_sidebar":true,"supports_join_request_queries":true}`

	var info BotFullInfo
	if err := json.Unmarshal([]byte(src), &info); err != nil {
		t.Fatal(err)
	}
	if info.ID != 42 || !info.IsBot || info.FirstName != "Bot" || info.Username != "foo" {
		t.Fatalf("%+v", info)
	}
	if len(info.ActiveUsernames) != 2 || info.ActiveUsernames[0] != "foo" || info.ActiveUsernames[1] != "foo_bot" {
		t.Fatalf("usernames %+v", info.ActiveUsernames)
	}
	if !info.SupportsInlineQueries || info.InlineQueryPlaceholder != "search" || !info.HasTopicSidebar {
		t.Fatalf("%+v", info)
	}
	if !info.NeedsInlineQueryLocation || !info.HasMainWebApp || !info.SupportsJoinRequestQueries {
		t.Fatalf("%+v", info)
	}

	var user User
	if err := json.Unmarshal([]byte(src), &user); err != nil {
		t.Fatal(err)
	}
	if user.ID != 42 || user.FirstName != "Bot" || user.Username != "foo" || !user.IsBot {
		t.Fatalf("%+v", user)
	}
}
