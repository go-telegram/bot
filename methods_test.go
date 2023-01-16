package bot

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type httpClient struct {
	resp      string
	reqFields map[string]string
	t         *testing.T
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	if c.reqFields != nil {
		for k, v := range c.reqFields {
			if v != req.FormValue(k) {
				c.t.Errorf("invalid field %q value %q, expect %q", k, req.FormValue(k), v)
				return nil, fmt.Errorf("invalid field %q value %q, expect %q", k, req.FormValue(k), v)
			}
		}
	}

	resp := &http.Response{}
	resp.Body = io.NopCloser(bytes.NewReader([]byte(`{"ok":true,"result":` + c.resp + `}`)))
	resp.StatusCode = http.StatusOK
	return resp, nil
}

func assertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func assertTrue(t *testing.T, v bool) {
	if !v {
		t.Errorf("unexpected not true")
	}
}

func assertEqualString(t *testing.T, v, e string) {
	if v != e {
		t.Errorf("strings is not equals, value %q, expect %q", v, e)
	}
}

func assertEqualInt(t *testing.T, v, e int) {
	if v != e {
		t.Errorf("ints is not equals, value %d, expect %d", v, e)
	}
}

func TestBot_Methods(t *testing.T) {
	t.Parallel()

	t.Run("SetWebhook", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"url":             "https://domain.com",
			"ip_address":      "1.2.3.4",
			"max_connections": "1",
		}}
		b := &Bot{client: c}
		resp, err := b.SetWebhook(context.Background(), &SetWebhookParams{
			URL:            "https://domain.com",
			IPAddress:      "1.2.3.4",
			MaxConnections: 1,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetWebhookInfo", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"url":"https://domain.com"}`}
		b := &Bot{client: c}
		resp, err := b.GetWebhookInfo(context.Background())
		assertNoErr(t, err)
		assertEqualString(t, resp.URL, "https://domain.com")
	})

	t.Run("DeleteWebhook", func(t *testing.T) {
		c := &httpClient{t: t, resp: "true", reqFields: map[string]string{
			"drop_pending_updates": "true",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteWebhook(context.Background(), &DeleteWebhookParams{
			DropPendingUpdates: true,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMe", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"username":"foo"}`}
		b := &Bot{client: c}
		resp, err := b.GetMe(context.Background())
		assertNoErr(t, err)
		assertEqualString(t, resp.Username, "foo")
	})

	t.Run("Logout", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`}
		b := &Bot{client: c}
		resp, err := b.Logout(context.Background())
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("Close", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`}
		b := &Bot{client: c}
		resp, err := b.Close(context.Background())
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SendMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"text": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendMessage(context.Background(), &SendMessageParams{
			Text: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("ForwardMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"message_id": "42",
		}}
		b := &Bot{client: c}
		resp, err := b.ForwardMessage(context.Background(), &ForwardMessageParams{
			MessageID: 42,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("CopyMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"message_id":12}`, reqFields: map[string]string{
			"message_id": "42",
		}}
		b := &Bot{client: c}
		resp, err := b.CopyMessage(context.Background(), &CopyMessageParams{
			MessageID: 42,
		})
		assertNoErr(t, err)
		assertEqualInt(t, 12, resp.ID)
	})

	t.Run("SendPhoto", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendPhoto(context.Background(), &SendPhotoParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendAudio", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendAudio(context.Background(), &SendAudioParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendDocument", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendDocument(context.Background(), &SendDocumentParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendVideo", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendVideo(context.Background(), &SendVideoParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendAnimation", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendAnimation(context.Background(), &SendAnimationParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendVoice", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"caption": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendVoice(context.Background(), &SendVoiceParams{
			Caption: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendVideoNote", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"thumb": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SendVideoNote(context.Background(), &SendVideoNoteParams{
			Thumb: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendMediaGroup", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[{"caption":"bar"}]`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendMediaGroup(context.Background(), &SendMediaGroupParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp[0].Caption)
	})

	t.Run("SendLocation", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendLocation(context.Background(), &SendLocationParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("EditMessageLiveLocation", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditMessageLiveLocation(context.Background(), &EditMessageLiveLocationParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("StopMessageLiveLocation", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.StopMessageLiveLocation(context.Background(), &StopMessageLiveLocationParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendVenue", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendVenue(context.Background(), &SendVenueParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendContact", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendContact(context.Background(), &SendContactParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendPoll", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendPoll(context.Background(), &SendPollParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendDice", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"bar"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendDice(context.Background(), &SendDiceParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Text)
	})

	t.Run("SendChatAction", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendChatAction(context.Background(), &SendChatActionParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetUserProfilePhotos", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"total_count":12}`, reqFields: map[string]string{
			"user_id": "234",
		}}
		b := &Bot{client: c}
		resp, err := b.GetUserProfilePhotos(context.Background(), &GetUserProfilePhotosParams{
			UserID: 234,
		})
		assertNoErr(t, err)
		assertEqualInt(t, 12, resp.TotalCount)
	})

	t.Run("GetFile", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"file_id":"123"}`, reqFields: map[string]string{
			"file_id": "234",
		}}
		b := &Bot{client: c}
		resp, err := b.GetFile(context.Background(), &GetFileParams{
			FileID: "234",
		})
		assertNoErr(t, err)
		assertEqualString(t, "123", resp.FileID)
	})

	t.Run("BanChatMember", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.BanChatMember(context.Background(), &BanChatMemberParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnbanChatMember", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnbanChatMember(context.Background(), &UnbanChatMemberParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("RestrictChatMember", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.RestrictChatMember(context.Background(), &RestrictChatMemberParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("PromoteChatMember", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.PromoteChatMember(context.Background(), &PromoteChatMemberParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetChatAdministratorCustomTitle", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatAdministratorCustomTitle(context.Background(), &SetChatAdministratorCustomTitleParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("BanChatSenderChat", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.BanChatSenderChat(context.Background(), &BanChatSenderChatParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnbanChatSenderChat", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnbanChatSenderChat(context.Background(), &UnbanChatSenderChatParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetChatPermissions", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatPermissions(context.Background(), &SetChatPermissionsParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("ExportChatInviteLink", func(t *testing.T) {
		c := &httpClient{t: t, resp: `"foo"`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.ExportChatInviteLink(context.Background(), &ExportChatInviteLinkParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp)
	})

	t.Run("CreateChatInviteLink", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CreateChatInviteLink(context.Background(), &CreateChatInviteLinkParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Name)
	})

	t.Run("EditChatInviteLink", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditChatInviteLink(context.Background(), &EditChatInviteLinkParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Name)
	})

	t.Run("RevokeChatInviteLink", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.RevokeChatInviteLink(context.Background(), &RevokeChatInviteLinkParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Name)
	})

}
