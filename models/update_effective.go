package models

// EffectiveUser returns the user associated with the update, or nil if none is
// available or u is nil. The result points into the update; it is not a copy.
// For messages it returns From (including channel posts), for callbacks and
// widget queries the user who triggered them, and for membership changes the
// initiating user.
// Business connections return their owner, subscriptions their subscriber,
// managed bot updates their User (not Bot), and boosts their source user.
// Anonymous actors are not inferred from chats or message contents.
func (u *Update) EffectiveUser() *User {
	if u == nil {
		return nil
	}
	switch {
	case u.Message != nil:
		return u.Message.From
	case u.EditedMessage != nil:
		return u.EditedMessage.From
	case u.ChannelPost != nil:
		return u.ChannelPost.From
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.From
	case u.BusinessConnection != nil:
		return &u.BusinessConnection.User
	case u.BusinessMessage != nil:
		return u.BusinessMessage.From
	case u.EditedBusinessMessage != nil:
		return u.EditedBusinessMessage.From
	case u.MessageReaction != nil:
		return u.MessageReaction.User
	case u.InlineQuery != nil:
		return u.InlineQuery.From
	case u.ChosenInlineResult != nil:
		return &u.ChosenInlineResult.From
	case u.CallbackQuery != nil:
		return &u.CallbackQuery.From
	case u.WidgetQuery != nil:
		return &u.WidgetQuery.From
	case u.SharedWidgetQuery != nil:
		return &u.SharedWidgetQuery.From
	case u.ShippingQuery != nil:
		return u.ShippingQuery.From
	case u.PreCheckoutQuery != nil:
		return u.PreCheckoutQuery.From
	case u.PurchasedPaidMedia != nil:
		return &u.PurchasedPaidMedia.From
	case u.PollAnswer != nil:
		return u.PollAnswer.User
	case u.ManagedBot != nil:
		return u.ManagedBot.User
	case u.GuestMessage != nil:
		return u.GuestMessage.From
	case u.MyChatMember != nil:
		return &u.MyChatMember.From
	case u.ChatMember != nil:
		return &u.ChatMember.From
	case u.ChatJoinRequest != nil:
		return &u.ChatJoinRequest.From
	case u.ChatBoost != nil:
		return effectiveBoostUser(&u.ChatBoost.Boost.Source)
	case u.RemovedChatBoost != nil:
		return effectiveBoostUser(&u.RemovedChatBoost.Source)
	case u.Subscription != nil:
		return &u.Subscription.User
	}
	return nil
}

func effectiveBoostUser(source *ChatBoostSource) *User {
	switch source.Source {
	case ChatBoostSourceTypePremium:
		if source.ChatBoostSourcePremium != nil {
			return &source.ChatBoostSourcePremium.User
		}
	case ChatBoostSourceTypeGiftCode:
		if source.ChatBoostSourceGiftCode != nil {
			return &source.ChatBoostSourceGiftCode.User
		}
	case ChatBoostSourceTypeGiveaway:
		// Giveaway.User is optional in the API but represented as a value here.
		if source.ChatBoostSourceGiveaway != nil && source.ChatBoostSourceGiveaway.User.ID != 0 {
			return &source.ChatBoostSourceGiveaway.User
		}
	}
	return nil
}

// EffectiveChat returns the chat where the event occurred, or nil if none is
// available or u is nil. The result points into the update; it is not a copy.
// Callback queries support both accessible and inaccessible messages; inline
// callbacks without a message have no chat. SenderChat, reaction ActorChat and
// PollAnswer.VoterChat identify actors, not the event's chat, and are not used.
// A chat is never synthesized from a user ID or BusinessConnection.UserChatID.
func (u *Update) EffectiveChat() *Chat {
	if u == nil {
		return nil
	}
	switch {
	case u.Message != nil:
		return &u.Message.Chat
	case u.EditedMessage != nil:
		return &u.EditedMessage.Chat
	case u.ChannelPost != nil:
		return &u.ChannelPost.Chat
	case u.EditedChannelPost != nil:
		return &u.EditedChannelPost.Chat
	case u.BusinessMessage != nil:
		return &u.BusinessMessage.Chat
	case u.EditedBusinessMessage != nil:
		return &u.EditedBusinessMessage.Chat
	case u.DeletedBusinessMessages != nil:
		return &u.DeletedBusinessMessages.Chat
	case u.MessageReaction != nil:
		return &u.MessageReaction.Chat
	case u.MessageReactionCount != nil:
		return &u.MessageReactionCount.Chat
	case u.CallbackQuery != nil:
		m := &u.CallbackQuery.Message
		switch m.Type {
		case MaybeInaccessibleMessageTypeMessage:
			if m.Message != nil {
				return &m.Message.Chat
			}
		case MaybeInaccessibleMessageTypeInaccessibleMessage:
			if m.InaccessibleMessage != nil {
				return &m.InaccessibleMessage.Chat
			}
		}
	case u.WidgetQuery != nil:
		return u.WidgetQuery.Chat
	case u.GuestMessage != nil:
		return &u.GuestMessage.Chat
	case u.MyChatMember != nil:
		return &u.MyChatMember.Chat
	case u.ChatMember != nil:
		return &u.ChatMember.Chat
	case u.ChatJoinRequest != nil:
		return &u.ChatJoinRequest.Chat
	case u.ChatBoost != nil:
		return &u.ChatBoost.Chat
	case u.RemovedChatBoost != nil:
		return &u.RemovedChatBoost.Chat
	case u.StoppedMessageGeneration != nil:
		return &u.StoppedMessageGeneration.Chat
	}
	return nil
}
