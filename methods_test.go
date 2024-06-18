package bot

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/go-telegram/bot/models"
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
			"thumbnail": `foo`,
		}}
		b := &Bot{client: c}
		resp, err := b.SendVideoNote(context.Background(), &SendVideoNoteParams{
			Thumbnail: &models.InputFileString{Data: "foo"},
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

	t.Run("ApproveChatJoinRequest", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.ApproveChatJoinRequest(context.Background(), &ApproveChatJoinRequestParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeclineChatJoinRequest", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeclineChatJoinRequest(context.Background(), &DeclineChatJoinRequestParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetChatPhoto", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatPhoto(context.Background(), &SetChatPhotoParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteChatPhoto", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteChatPhoto(context.Background(), &DeleteChatPhotoParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetChatTitle", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatTitle(context.Background(), &SetChatTitleParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetChatDescription", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatDescription(context.Background(), &SetChatDescriptionParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("PinChatMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.PinChatMessage(context.Background(), &PinChatMessageParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnpinChatMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnpinChatMessage(context.Background(), &UnpinChatMessageParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnpinAllChatMessages", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnpinAllChatMessages(context.Background(), &UnpinAllChatMessagesParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("LeaveChat", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.LeaveChat(context.Background(), &LeaveChatParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetChat", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"title":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetChat(context.Background(), &GetChatParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Title)
	})

	t.Run("GetChatAdministrators", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[]`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetChatAdministrators(context.Background(), &GetChatAdministratorsParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualInt(t, 0, len(resp))
	})

	t.Run("GetChatMemberCount", func(t *testing.T) {
		c := &httpClient{t: t, resp: `42`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetChatMemberCount(context.Background(), &GetChatMemberCountParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualInt(t, 42, resp)
	})

	t.Run("GetChatMember", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"status":"administrator"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetChatMember(context.Background(), &GetChatMemberParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, string(models.ChatMemberTypeAdministrator), string(resp.Type))
	})

	t.Run("SetChatStickerSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatStickerSet(context.Background(), &SetChatStickerSetParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetForumTopicIconStickers", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[]`, reqFields: map[string]string{}}
		b := &Bot{client: c}
		resp, err := b.GetForumTopicIconStickers(context.Background())
		assertNoErr(t, err)
		assertEqualInt(t, 0, len(resp))
	})

	t.Run("CreateForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CreateForumTopic(context.Background(), &CreateForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Name)
	})

	t.Run("EditForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditForumTopic(context.Background(), &EditForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("CloseForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CloseForumTopic(context.Background(), &CloseForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("ReopenForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.ReopenForumTopic(context.Background(), &ReopenForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnpinAllForumTopicMessages", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnpinAllForumTopicMessages(context.Background(), &UnpinAllForumTopicMessagesParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("EditGeneralForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditGeneralForumTopic(context.Background(), &EditGeneralForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("CloseGeneralForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CloseGeneralForumTopic(context.Background(), &CloseGeneralForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("ReopenGeneralForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.ReopenGeneralForumTopic(context.Background(), &ReopenGeneralForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("HideGeneralForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.HideGeneralForumTopic(context.Background(), &HideGeneralForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("UnhideGeneralForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UnhideGeneralForumTopic(context.Background(), &UnhideGeneralForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteForumTopic", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteForumTopic(context.Background(), &DeleteForumTopicParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteChatStickerSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteChatStickerSet(context.Background(), &DeleteChatStickerSetParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("AnswerCallbackQuery", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"text": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.AnswerCallbackQuery(context.Background(), &AnswerCallbackQueryParams{
			Text: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetMyCommands", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SetMyCommands(context.Background(), &SetMyCommandsParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteMyCommands", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteMyCommands(context.Background(), &DeleteMyCommandsParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMyCommands", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[]`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.GetMyCommands(context.Background(), &GetMyCommandsParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertEqualInt(t, 0, len(resp))
	})

	t.Run("SetMyName", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"name":          "bar",
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SetMyName(context.Background(), &SetMyNameParams{
			Name:         "bar",
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMyName", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"bar"}`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.GetMyName(context.Background(), &GetMyNameParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "bar", resp.Name)
	})

	t.Run("SetMyDescription", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SetMyDescription(context.Background(), &SetMyDescriptionParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMyDescription", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"description":"foo"}`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.GetMyDescription(context.Background(), &GetMyDescriptionParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Description)
	})

	t.Run("SetMyShortDescription", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.SetMyShortDescription(context.Background(), &SetMyShortDescriptionParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMyShortDescription", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"short_description":"foo"}`, reqFields: map[string]string{
			"language_code": "foo",
		}}
		b := &Bot{client: c}
		resp, err := b.GetMyShortDescription(context.Background(), &GetMyShortDescriptionParams{
			LanguageCode: "foo",
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.ShortDescription)
	})

	t.Run("SetChatMenuButton", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetChatMenuButton(context.Background(), &SetChatMenuButtonParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetChatMenuButton", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"type":"commands"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetChatMenuButton(context.Background(), &GetChatMenuButtonParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp.Type == models.MenuButtonTypeCommands)
	})

	t.Run("SetMyDefaultAdministratorRights", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"for_channels": "true",
		}}
		b := &Bot{client: c}
		resp, err := b.SetMyDefaultAdministratorRights(context.Background(), &SetMyDefaultAdministratorRightsParams{
			ForChannels: true,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("GetMyDefaultAdministratorRights", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"can_edit_messages":true}`, reqFields: map[string]string{
			"for_channels": "true",
		}}
		b := &Bot{client: c}
		resp, err := b.GetMyDefaultAdministratorRights(context.Background(), &GetMyDefaultAdministratorRightsParams{
			ForChannels: true,
		})
		assertNoErr(t, err)
		assertTrue(t, resp.CanEditMessages)
	})

	t.Run("EditMessageText", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditMessageText(context.Background(), &EditMessageTextParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("EditMessageCaption", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditMessageCaption(context.Background(), &EditMessageCaptionParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("EditMessageMedia", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditMessageMedia(context.Background(), &EditMessageMediaParams{
			ChatID: 123,
			Media:  &models.InputMediaPhoto{},
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("EditMessageReplyMarkup", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.EditMessageReplyMarkup(context.Background(), &EditMessageReplyMarkupParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("StopPoll", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"id":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.StopPoll(context.Background(), &StopPollParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.ID)
	})

	t.Run("DeleteMessage", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteMessage(context.Background(), &DeleteMessageParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SendSticker", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendSticker(context.Background(), &SendStickerParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("GetStickerSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"name":"foo"}`, reqFields: map[string]string{
			"name": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetStickerSet(context.Background(), &GetStickerSetParams{
			Name: "123",
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Name)
	})

	t.Run("GetCustomEmojiStickers", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[]`, reqFields: map[string]string{}}
		b := &Bot{client: c}
		resp, err := b.GetCustomEmojiStickers(context.Background(), &GetCustomEmojiStickersParams{})
		assertNoErr(t, err)
		assertEqualInt(t, 0, len(resp))
	})

	t.Run("UploadStickerFile", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"file_id":"foo"}`, reqFields: map[string]string{
			"user_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.UploadStickerFile(context.Background(), &UploadStickerFileParams{
			UserID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.FileID)
	})

	t.Run("CreateNewStickerSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"user_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CreateNewStickerSet(context.Background(), &CreateNewStickerSetParams{
			UserID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("AddStickerToSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"user_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.AddStickerToSet(context.Background(), &AddStickerToSetParams{
			UserID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerPositionInSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"sticker": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerPositionInSet(context.Background(), &SetStickerPositionInSetParams{
			Sticker: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteStickerFromSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"sticker": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteStickerFromSet(context.Background(), &DeleteStickerFromSetParams{
			Sticker: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerEmojiList", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"sticker": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerEmojiList(context.Background(), &SetStickerEmojiListParams{
			Sticker: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerKeywords", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"sticker": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerKeywords(context.Background(), &SetStickerKeywordsParams{
			Sticker: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerMaskPosition", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"sticker": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerMaskPosition(context.Background(), &SetStickerMaskPositionParams{
			Sticker: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerSetTitle", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"title": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerSetTitle(context.Background(), &SetStickerSetTitleParams{
			Title: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetStickerSetThumbnail", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"name": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetStickerSetThumbnail(context.Background(), &SetStickerSetThumbnailParams{
			Name: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetCustomEmojiStickerSetThumbnail", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"name": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetCustomEmojiStickerSetThumbnail(context.Background(), &SetCustomEmojiStickerSetThumbnailParams{
			Name: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("DeleteStickerSet", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"name": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.DeleteStickerSet(context.Background(), &DeleteStickerSetParams{
			Name: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("AnswerInlineQuery", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"inline_query_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.AnswerInlineQuery(context.Background(), &AnswerInlineQueryParams{
			InlineQueryID: "123",
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("AnswerWebAppQuery", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"inline_message_id":"foo"}`, reqFields: map[string]string{
			"web_app_query_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.AnswerWebAppQuery(context.Background(), &AnswerWebAppQueryParams{
			WebAppQueryID: "123",
			Result:        &models.InlineQueryResultArticle{},
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.InlineMessageID)
	})

	t.Run("SendInvoice", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendInvoice(context.Background(), &SendInvoiceParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("CreateInvoiceLink", func(t *testing.T) {
		c := &httpClient{t: t, resp: `"foo"`, reqFields: map[string]string{
			"title": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.CreateInvoiceLink(context.Background(), &CreateInvoiceLinkParams{
			Title: "123",
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp)
	})

	t.Run("AnswerShippingQuery", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"ok": "true",
		}}
		b := &Bot{client: c}
		resp, err := b.AnswerShippingQuery(context.Background(), &AnswerShippingQueryParams{
			OK: true,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("AnswerPreCheckoutQuery", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"ok": "true",
		}}
		b := &Bot{client: c}
		resp, err := b.AnswerPreCheckoutQuery(context.Background(), &AnswerPreCheckoutQueryParams{
			OK: true,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SetPassportDataErrors", func(t *testing.T) {
		c := &httpClient{t: t, resp: `true`, reqFields: map[string]string{
			"user_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetPassportDataErrors(context.Background(), &SetPassportDataErrorsParams{
			UserID: 123,
		})
		assertNoErr(t, err)
		assertTrue(t, resp)
	})

	t.Run("SendGame", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SendGame(context.Background(), &SendGameParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("SetGameScore", func(t *testing.T) {
		c := &httpClient{t: t, resp: `{"text":"foo"}`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.SetGameScore(context.Background(), &SetGameScoreParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualString(t, "foo", resp.Text)
	})

	t.Run("GetGameHighScores", func(t *testing.T) {
		c := &httpClient{t: t, resp: `[{"score":42}]`, reqFields: map[string]string{
			"chat_id": "123",
		}}
		b := &Bot{client: c}
		resp, err := b.GetGameHighScores(context.Background(), &GetGameHighScoresParams{
			ChatID: 123,
		})
		assertNoErr(t, err)
		assertEqualInt(t, 42, resp[0].Score)
	})

}
