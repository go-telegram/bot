package models

// Update https://core.telegram.org/bots/api#update
type Update struct {
	ID                 int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}
