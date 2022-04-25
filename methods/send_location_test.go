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

func TestSendLocationParams_Validate(t *testing.T) {
	type fields struct {
		ChatID    string
		Latitude  float64
		Longitude float64
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
				ChatID:    "",
				Latitude:  0,
				Longitude: 0,
			},
			wantErr:  true,
			errValue: "empty chat id",
		},
		{
			name: "empty latitude",
			fields: fields{
				ChatID:    "foo",
				Latitude:  0,
				Longitude: 0,
			},
			wantErr:  true,
			errValue: "empty latitude",
		},
		{
			name: "empty longitude",
			fields: fields{
				ChatID:    "foo",
				Latitude:  10,
				Longitude: 0,
			},
			wantErr:  true,
			errValue: "empty longitude",
		},
		{
			name: "ok",
			fields: fields{
				ChatID:    "foo",
				Latitude:  10,
				Longitude: 10,
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SendLocationParams{
				ChatID:    tt.fields.ChatID,
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
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

func TestSendLocation(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true,"result":{"message_id":12}}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := SendLocation(context.Background(), b, &SendLocationParams{
		ChatID:    "foo",
		Latitude:  10,
		Longitude: 10,
	})
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	expect := &models.Message{
		ID: 12,
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Expected: %v, got: %v", expect, res)
	}
}
