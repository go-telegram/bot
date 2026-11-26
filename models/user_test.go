package models

import (
	"encoding/json"
	"strings"
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

func TestBotFullInfo_MarshalRequiredFalse(t *testing.T) {
	out, err := json.Marshal(BotFullInfo{ID: 1, IsBot: true, FirstName: "Bot"})
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	if !strings.Contains(got, `"can_join_groups":false`) {
		t.Fatalf("required false dropped: %s", got)
	}
	for _, optional := range []string{`"username"`, `"active_usernames"`, `"inline_query_placeholder"`} {
		if strings.Contains(got, optional) {
			t.Fatalf("optional %s present in %s", optional, got)
		}
	}
}

func TestBotFullInfo_SupportsJoinRequestQueries(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{
			name: "with supports_join_request_queries",
			src:  `{"id":1,"is_bot":true,"first_name":"Bot","supports_join_request_queries":true}`,
			want: true,
		},
		{
			name: "without supports_join_request_queries",
			src:  `{"id":1,"is_bot":true,"first_name":"Bot"}`,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var info BotFullInfo
			if err := json.Unmarshal([]byte(tt.src), &info); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if info.SupportsJoinRequestQueries != tt.want {
				t.Fatalf("supports_join_request_queries = %v, want %v", info.SupportsJoinRequestQueries, tt.want)
			}
		})
	}
}

func TestBotFullInfo_SupportsInlineQueries(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{
			name: "with supports_inline_queries",
			src:  `{"id":1,"is_bot":true,"first_name":"Bot","supports_inline_queries":true}`,
			want: true,
		},
		{
			name: "without supports_inline_queries",
			src:  `{"id":1,"is_bot":true,"first_name":"Bot"}`,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var info BotFullInfo
			if err := json.Unmarshal([]byte(tt.src), &info); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if info.SupportsInlineQueries != tt.want {
				t.Fatalf("supports_inline_queries = %v, want %v", info.SupportsInlineQueries, tt.want)
			}
		})
	}
}
