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
