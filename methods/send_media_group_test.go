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

func TestSendMediaGroupParams_Validate(t *testing.T) {
	type fields struct {
		ChatID string
		Media  []InputMedia
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
				ChatID: "",
				Media:  []InputMedia{},
			},
			wantErr:  true,
			errValue: "empty chat id",
		},
		{
			name: "empty media",
			fields: fields{
				ChatID: "foo",
				Media:  []InputMedia{},
			},
			wantErr:  true,
			errValue: "empty media",
		},
		{
			name: "ok",
			fields: fields{
				ChatID: "foo",
				Media:  []InputMedia{InputMediaPhoto{}},
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SendMediaGroupParams{
				ChatID: tt.fields.ChatID,
				Media:  tt.fields.Media,
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

func TestSendMediaGroup(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true,"result":[{"message_id":12}]}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := SendMediaGroup(context.Background(), b, &SendMediaGroupParams{
		ChatID: "foo",
		Media:  []InputMedia{InputMediaPhoto{}},
	})
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	expect := []*models.Message{{
		ID: 12,
	}}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Expected: %v, got: %v", expect, res)
	}
}
