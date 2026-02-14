package bot

import "github.com/go-telegram/bot/models"

func ExtractUserID(upd *models.Update) (int64, bool) {
	if upd != nil {
		switch {
		case upd.Message != nil && upd.Message.From != nil:
			return upd.Message.From.ID, true
		case upd.EditedMessage != nil && upd.EditedMessage.From != nil:
			return upd.EditedMessage.From.ID, true
		case upd.ChannelPost != nil && upd.ChannelPost.From != nil:
			return upd.ChannelPost.From.ID, true
		case upd.EditedChannelPost != nil && upd.EditedChannelPost.From != nil:
			return upd.EditedChannelPost.From.ID, true
		case upd.BusinessConnection != nil:
			return upd.BusinessConnection.User.ID, true
		case upd.BusinessMessage != nil && upd.BusinessMessage.From != nil:
			return upd.BusinessMessage.From.ID, true
		case upd.EditedBusinessMessage != nil && upd.EditedBusinessMessage.From != nil:
			return upd.EditedBusinessMessage.From.ID, true
		case upd.MessageReaction != nil && upd.MessageReaction.User != nil:
			return upd.MessageReaction.User.ID, true
		case upd.InlineQuery != nil && upd.InlineQuery.From != nil:
			return upd.InlineQuery.From.ID, true
		case upd.ChosenInlineResult != nil:
			return upd.ChosenInlineResult.From.ID, true
		case upd.CallbackQuery != nil:
			return upd.CallbackQuery.From.ID, true
		case upd.ShippingQuery != nil && upd.ShippingQuery.From != nil:
			return upd.ShippingQuery.From.ID, true
		case upd.PreCheckoutQuery != nil && upd.PreCheckoutQuery.From != nil:
			return upd.PreCheckoutQuery.From.ID, true
		case upd.PurchasedPaidMedia != nil:
			return upd.PurchasedPaidMedia.From.ID, true
		case upd.PollAnswer != nil && upd.PollAnswer.User != nil:
			return upd.PollAnswer.User.ID, true
		case upd.MyChatMember != nil:
			return upd.MyChatMember.From.ID, true
		case upd.ChatMember != nil:
			return upd.ChatMember.From.ID, true
		case upd.ChatJoinRequest != nil:
			return upd.ChatJoinRequest.From.ID, true
		}
	}

	return 0, false
}
