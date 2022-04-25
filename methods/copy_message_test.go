package methods

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestCopyMessageParams_Validate(t *testing.T) {
	type fields struct {
		ChatID     string
		FromChatID string
		MessageID  int
	}
	tests := []struct {
		name     string
		fields   fields
		wantErr  bool
		errValue string
	}{
		{
			name: "empty chat id",
			fields: fields{
				ChatID:     "",
				FromChatID: "foo",
				MessageID:  42,
			},
			wantErr:  true,
			errValue: "empty chat id",
		},
		{
			name: "empty from chat id",
			fields: fields{
				ChatID:     "foo",
				FromChatID: "",
				MessageID:  42,
			},
			wantErr:  true,
			errValue: "empty from chat id",
		},
		{
			name: "empty message id",
			fields: fields{
				ChatID:     "foo",
				FromChatID: "bar",
				MessageID:  0,
			},
			wantErr:  true,
			errValue: "empty message id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CopyMessageParams{
				ChatID:     tt.fields.ChatID,
				FromChatID: tt.fields.FromChatID,
				MessageID:  tt.fields.MessageID,
			}
			err := p.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !tt.wantErr {
				return
			}
			if (err != nil) && tt.wantErr {
				if err.Error() != tt.errValue {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.errValue)
				}
			}
		})
	}
}

func TestCopyMessage(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true,"result":{"message_id":12}}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := CopyMessage(context.Background(), b, &CopyMessageParams{
		ChatID:     "foo",
		FromChatID: "bar",
		MessageID:  42,
	})
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	expect := &models.MessageID{
		ID: 12,
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Expected: %v, got: %v", expect, res)
	}
}
