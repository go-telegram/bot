package methods

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TestSendAudioParams_Validate(t *testing.T) {
	type fields struct {
		ChatID string
		Audio  models.InputFile
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
				Audio:  &models.InputFileString{Data: "foo"},
			},
			wantErr:  true,
			errValue: "empty chat id",
		},
		{
			name: "empty audio",
			fields: fields{
				ChatID: "foo",
				Audio:  nil,
			},
			wantErr:  true,
			errValue: "empty audio",
		},
		{
			name: "ok",
			fields: fields{
				ChatID: "foo",
				Audio:  &models.InputFileString{Data: "ok"},
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SendAudioParams{
				ChatID: tt.fields.ChatID,
				Audio:  tt.fields.Audio,
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

func TestSendAudio(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true,"result":{"message_id":12}}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := SendAudio(context.Background(), b, &SendAudioParams{
		ChatID: "foo",
		Audio:  &models.InputFileString{Data: "bar"},
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
