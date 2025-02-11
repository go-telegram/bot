package models

// Update https://core.telegram.org/bots/api#update
type Update struct {
	ID                      int64                        `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

// allowed_updates https://core.telegram.org/bots/api#update
const (
	AllowedUpdateMessage                 string = "message"
	AllowedUpdateEditedMessage           string = "edited_message"
	AllowedUpdateChannelPost             string = "channel_post"
	AllowedUpdateEditedChannelPost       string = "edited_channel_post"
	AllowedUpdateBusinessConnection      string = "business_connection"
	AllowedUpdateBusinessMessage         string = "business_message"
	AllowedUpdateEditedBusinessMessage   string = "edited_business_message"
	AllowedUpdateDeletedBusinessMessages string = "deleted_business_messages"
	AllowedUpdateMessageReaction         string = "message_reaction"
	AllowedUpdateMessageReactionCount    string = "message_reaction_count"
	AllowedUpdateInlineQuery             string = "inline_query"
	AllowedUpdateChosenInlineResult      string = "chosen_inline_result"
	AllowedUpdateCallbackQuery           string = "callback_query"
	AllowedUpdateShippingQuery           string = "shipping_query"
	AllowedUpdatePreCheckoutQuery        string = "pre_checkout_query"
	AllowedUpdatePurchasedPaidMedia      string = "purchased_paid_media"
	AllowedUpdatePoll                    string = "poll"
	AllowedUpdatePollAnswer              string = "poll_answer"
	AllowedUpdateMyChatMember            string = "my_chat_member"
	AllowedUpdateChatMember              string = "chat_member"
	AllowedUpdateChatJoinRequest         string = "chat_join_request"
	AllowedUpdateChatBoost               string = "chat_boost"
	AllowedUpdateRemovedChatBoost        string = "removed_chat_boost"
)

func (u *Update) FromUser() *User {
	if u.Message != nil {
		return u.Message.From
	} else if u.EditedMessage != nil {
		return u.EditedMessage.From
	} else if u.ChannelPost != nil {
		return u.ChannelPost.From
	} else if u.EditedChannelPost != nil {
		return u.EditedChannelPost.From
	} else if u.BusinessConnection != nil {
		return &u.BusinessConnection.User
	} else if u.BusinessMessage != nil {
		return u.BusinessMessage.From
	} else if u.EditedBusinessMessage != nil {
		return u.EditedBusinessMessage.From
	} else if u.MessageReaction != nil {
		return u.EditedMessage.From
	} else if u.InlineQuery != nil {
		return u.InlineQuery.From
	} else if u.ChosenInlineResult != nil {
		return &u.ChosenInlineResult.From
	} else if u.CallbackQuery != nil {
		return &u.CallbackQuery.From
	} else if u.ShippingQuery != nil {
		return u.ShippingQuery.From
	} else if u.PreCheckoutQuery != nil {
		return u.PreCheckoutQuery.From
	} else if u.PurchasedPaidMedia != nil {
		return &u.PurchasedPaidMedia.From
	} else if u.PollAnswer != nil {
		return u.PollAnswer.User
	} else if u.MyChatMember != nil {
		return &u.MyChatMember.From
	} else if u.ChatMember != nil {
		return &u.ChatMember.From
	} else if u.ChatJoinRequest != nil {
		return &u.ChatJoinRequest.From
	} else if u.ChatBoost != nil {
		return u.ChatBoost.Boost.Source.FromUser()
	} else if u.RemovedChatBoost != nil {
		return u.RemovedChatBoost.Source.FromUser()
	} else {
		// Telegram can add new update types, so there is possibility of forgetting to add the statement
		// or update simply could miss this data

		// Poll, MessageReactionCount, DeletedBusinessMessages are skipped, they don't have User field
		return nil
	}
}

func (u *Update) FromChat() *Chat {
	if u.Message != nil {
		return &u.Message.Chat
	} else if u.EditedMessage != nil {
		return &u.EditedMessage.Chat
	} else if u.ChannelPost != nil {
		return &u.ChannelPost.Chat
	} else if u.EditedChannelPost != nil {
		return &u.EditedChannelPost.Chat
	} else if u.BusinessMessage != nil {
		return &u.BusinessMessage.Chat
	} else if u.EditedBusinessMessage != nil {
		return &u.EditedBusinessMessage.Chat
	} else if u.MessageReaction != nil {
		return &u.EditedMessage.Chat
	} else if u.CallbackQuery != nil {
		return u.CallbackQuery.Message.FromChat()
	} else if u.PollAnswer != nil {
		return u.PollAnswer.VoterChat
	} else if u.MyChatMember != nil {
		return &u.MyChatMember.Chat
	} else if u.ChatMember != nil {
		return &u.ChatMember.Chat
	} else if u.ChatJoinRequest != nil {
		return &u.ChatJoinRequest.Chat
	} else {
		// Telegram can add new update types, so there is possibility of forgetting to add the statement
		// or update simply could miss this data

		// BusinessConnection, InlineQuery, ChosenInlineResult
		// ShippingQuery, PreCheckoutQuery, PurchasedPaidMedia
		// ChatBoost, RemovedChatBoost
		// are skipped, they don't have User field

		// Poll, MessageReactionCount, DeletedBusinessMessages are skipped, they don't have User field
		return nil
	}
}
