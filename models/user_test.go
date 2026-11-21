package models

import (
	"encoding/json"
	"testing"
)

// TestUser_SupportsJoinRequestQueries verifies supports_join_request_queries decodes when present and stays false when absent.
func TestUser_SupportsJoinRequestQueries(t *testing.T) {
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
			var u User
			if err := json.Unmarshal([]byte(tt.src), &u); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if u.SupportsJoinRequestQueries != tt.want {
				t.Fatalf("supports_join_request_queries = %v, want %v", u.SupportsJoinRequestQueries, tt.want)
			}
		})
	}
}
