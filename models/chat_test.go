package models

import (
	"encoding/json"
	"testing"
)

// TestChatJoinRequest_QueryID verifies query_id decodes when present and stays empty when absent.
func TestChatJoinRequest_QueryID(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "with query_id",
			src:  `{"chat":{"id":1,"type":"supergroup"},"from":{"id":2,"is_bot":false,"first_name":"A"},"user_chat_id":2,"date":1,"query_id":"q1"}`,
			want: "q1",
		},
		{
			name: "without query_id",
			src:  `{"chat":{"id":1,"type":"supergroup"},"from":{"id":2,"is_bot":false,"first_name":"A"},"user_chat_id":2,"date":1}`,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r ChatJoinRequest
			if err := json.Unmarshal([]byte(tt.src), &r); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if r.QueryID != tt.want {
				t.Fatalf("query_id = %q, want %q", r.QueryID, tt.want)
			}
		})
	}
}

// TestChatFullInfo_GuardBot verifies guard_bot decodes when present and stays nil when absent.
func TestChatFullInfo_GuardBot(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want *User
	}{
		{
			name: "with guard_bot",
			src:  `{"id":1,"type":"supergroup","accent_color_id":0,"max_reaction_count":11,"guard_bot":{"id":42,"is_bot":true,"first_name":"Guard","username":"guard_bot"}}`,
			want: &User{ID: 42, IsBot: true, FirstName: "Guard", Username: "guard_bot"},
		},
		{
			name: "without guard_bot",
			src:  `{"id":1,"type":"supergroup","accent_color_id":0,"max_reaction_count":11}`,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c ChatFullInfo
			if err := json.Unmarshal([]byte(tt.src), &c); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if tt.want == nil {
				if c.GuardBot != nil {
					t.Fatalf("guard_bot = %+v, want nil", c.GuardBot)
				}
				return
			}
			if c.GuardBot == nil {
				t.Fatal("guard_bot not decoded")
			}
			if *c.GuardBot != *tt.want {
				t.Fatalf("guard_bot = %+v, want %+v", *c.GuardBot, *tt.want)
			}
		})
	}
}
