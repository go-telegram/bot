package bot

import (
	"context"

	"github.com/go-telegram/bot/models"
)

// SetWebhook https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(ctx context.Context, params *SetWebhookParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setWebhook", params, &result)
	return result, err
}

// GetWebhookInfo https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error) {
	result := &models.WebhookInfo{}
	err := b.rawRequest(ctx, "getWebhookInfo", nil, &result)
	return result, err
}

// DeleteWebhook https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(ctx context.Context, params *DeleteWebhookParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteWebhook", params, &result)
	return result, err
}

// GetMe https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe(ctx context.Context) (*models.User, error) {
	result := &models.User{}
	err := b.rawRequest(ctx, "getMe", nil, result)
	return result, err
}

// Logout https://core.telegram.org/bots/api#logout
func (b *Bot) Logout(ctx context.Context) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "logout", nil, &result)
	return result, err
}

// Close https://core.telegram.org/bots/api#close
func (b *Bot) Close(ctx context.Context) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "close", nil, &result)
	return result, err
}

// SendMessage https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(ctx context.Context, params *SendMessageParams) (*models.Message, error) {
	mes := &models.Message{}
	err := b.rawRequest(ctx, "sendMessage", params, mes)
	return mes, err
}

// ForwardMessage https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(ctx context.Context, params *ForwardMessageParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "forwardMessage", params, result)
	return result, err
}

// ForwardMessages https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(ctx context.Context, params *ForwardMessagesParams) ([]models.MessageID, error) {
	var result []models.MessageID
	err := b.rawRequest(ctx, "forwardMessages", params, &result)
	return result, err
}

// CopyMessage https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(ctx context.Context, params *CopyMessageParams) (*models.MessageID, error) {
	result := &models.MessageID{}
	err := b.rawRequest(ctx, "copyMessage", params, result)
	return result, err
}

// CopyMessages https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(ctx context.Context, params *CopyMessagesParams) ([]models.MessageID, error) {
	var result []models.MessageID
	err := b.rawRequest(ctx, "copyMessages", params, &result)
	return result, err
}

// SendPhoto https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(ctx context.Context, params *SendPhotoParams) (*models.Message, error) {
	mes := &models.Message{}
	err := b.rawRequest(ctx, "sendPhoto", params, mes)
	return mes, err
}

// SendAudio https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(ctx context.Context, params *SendAudioParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendAudio", params, result)
	return result, err
}

// SendDocument https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(ctx context.Context, params *SendDocumentParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendDocument", params, result)
	return result, err
}

// SendVideo https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(ctx context.Context, params *SendVideoParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendVideo", params, result)
	return result, err
}

// SendAnimation https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(ctx context.Context, params *SendAnimationParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendAnimation", params, result)
	return result, err
}

// SendVoice https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(ctx context.Context, params *SendVoiceParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendVoice", params, result)
	return result, err
}

// SendVideoNote https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(ctx context.Context, params *SendVideoNoteParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendVideoNote", params, result)
	return result, err
}

// SendPaidMedia https://core.telegram.org/bots/api#sendpaidmedia
func (b *Bot) SendPaidMedia(ctx context.Context, params *SendPaidMediaParams) (*models.Message, error) {
	var result models.Message
	err := b.rawRequest(ctx, "sendPaidMedia", params, &result)
	return &result, err
}

// SendMediaGroup https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(ctx context.Context, params *SendMediaGroupParams) ([]*models.Message, error) {
	var result []*models.Message
	err := b.rawRequest(ctx, "sendMediaGroup", params, &result)
	return result, err
}

// SendLocation https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(ctx context.Context, params *SendLocationParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendLocation", params, result)
	return result, err
}

// EditMessageLiveLocation https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(ctx context.Context, params *EditMessageLiveLocationParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "editMessageLiveLocation", params, result)
	return result, err
}

// StopMessageLiveLocation https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(ctx context.Context, params *StopMessageLiveLocationParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "stopMessageLiveLocation", params, result)
	return result, err
}

// SendVenue https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(ctx context.Context, params *SendVenueParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendVenue", params, result)
	return result, err
}

// SendContact https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(ctx context.Context, params *SendContactParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendContact", params, result)
	return result, err
}

// SendPoll https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(ctx context.Context, params *SendPollParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendPoll", params, result)
	return result, err
}

// SendDice https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(ctx context.Context, params *SendDiceParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendDice", params, &result)
	return result, err
}

// SendChatAction https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(ctx context.Context, params *SendChatActionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "sendChatAction", params, &result)
	return result, err
}

// SetMessageReaction https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(ctx context.Context, params *SetMessageReactionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMessageReaction", params, &result)
	return result, err
}

// GetUserProfilePhotos https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(ctx context.Context, params *GetUserProfilePhotosParams) (*models.UserProfilePhotos, error) {
	result := &models.UserProfilePhotos{}
	err := b.rawRequest(ctx, "getUserProfilePhotos", params, result)
	return result, err
}

// SetUserEmojiStatus https://core.telegram.org/bots/api#setuseremojistatus
func (b *Bot) SetUserEmojiStatus(ctx context.Context, params *SetUserEmojiStatusParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setUserEmojiStatus", params, &result)
	return result, err
}

// GetFile https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(ctx context.Context, params *GetFileParams) (*models.File, error) {
	result := &models.File{}
	err := b.rawRequest(ctx, "getFile", params, result)
	return result, err
}

// BanChatMember https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(ctx context.Context, params *BanChatMemberParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "banChatMember", params, &result)
	return result, err
}

// UnbanChatMember https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(ctx context.Context, params *UnbanChatMemberParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unbanChatMember", params, &result)
	return result, err
}

// RestrictChatMember https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(ctx context.Context, params *RestrictChatMemberParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "restrictChatMember", params, &result)
	return result, err
}

// PromoteChatMember https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(ctx context.Context, params *PromoteChatMemberParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "promoteChatMember", params, &result)
	return result, err
}

// SetChatAdministratorCustomTitle https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(ctx context.Context, params *SetChatAdministratorCustomTitleParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatAdministratorCustomTitle", params, &result)
	return result, err
}

// BanChatSenderChat https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(ctx context.Context, params *BanChatSenderChatParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "banChatSenderChat", params, &result)
	return result, err
}

// UnbanChatSenderChat https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(ctx context.Context, params *UnbanChatSenderChatParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unbanChatSenderChat", params, &result)
	return result, err
}

// SetChatPermissions https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(ctx context.Context, params *SetChatPermissionsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatPermissions", params, &result)
	return result, err
}

// ExportChatInviteLink https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(ctx context.Context, params *ExportChatInviteLinkParams) (string, error) {
	var result string
	err := b.rawRequest(ctx, "exportChatInviteLink", params, &result)
	return result, err
}

// CreateChatInviteLink https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(ctx context.Context, params *CreateChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}
	err := b.rawRequest(ctx, "createChatInviteLink", params, &result)
	return result, err
}

// EditChatInviteLink https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(ctx context.Context, params *EditChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}
	err := b.rawRequest(ctx, "editChatInviteLink", params, &result)
	return result, err
}

// CreateChatSubscriptionInviteLink https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (b *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, params *CreateChatSubscriptionInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}
	err := b.rawRequest(ctx, "createChatSubscriptionInviteLink", params, &result)
	return result, err
}

// EditChatSubscriptionInviteLink https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (b *Bot) EditChatSubscriptionInviteLink(ctx context.Context, params *EditChatSubscriptionInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}
	err := b.rawRequest(ctx, "editChatSubscriptionInviteLink", params, &result)
	return result, err
}

// RevokeChatInviteLink https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(ctx context.Context, params *RevokeChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}
	err := b.rawRequest(ctx, "revokeChatInviteLink", params, &result)
	return result, err
}

// ApproveChatJoinRequest https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(ctx context.Context, params *ApproveChatJoinRequestParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "approveChatJoinRequest", params, &result)
	return result, err
}

// DeclineChatJoinRequest https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(ctx context.Context, params *DeclineChatJoinRequestParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "declineChatJoinRequest", params, &result)
	return result, err
}

// SetChatPhoto https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(ctx context.Context, params *SetChatPhotoParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatPhoto", params, &result)
	return result, err
}

// DeleteChatPhoto https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(ctx context.Context, params *DeleteChatPhotoParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteChatPhoto", params, &result)
	return result, err
}

// SetChatTitle https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(ctx context.Context, params *SetChatTitleParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatTitle", params, &result)
	return result, err
}

// SetChatDescription https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(ctx context.Context, params *SetChatDescriptionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatDescription", params, &result)
	return result, err
}

// PinChatMessage https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(ctx context.Context, params *PinChatMessageParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "pinChatMessage", params, &result)
	return result, err
}

// UnpinChatMessage https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(ctx context.Context, params *UnpinChatMessageParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unpinChatMessage", params, &result)
	return result, err
}

// UnpinAllChatMessages https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(ctx context.Context, params *UnpinAllChatMessagesParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unpinAllChatMessages", params, &result)
	return result, err
}

// LeaveChat https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(ctx context.Context, params *LeaveChatParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "leaveChat", params, &result)
	return result, err
}

// GetChat https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(ctx context.Context, params *GetChatParams) (*models.ChatFullInfo, error) {
	var result *models.ChatFullInfo
	err := b.rawRequest(ctx, "getChat", params, &result)
	return result, err
}

// GetChatAdministrators https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(ctx context.Context, params *GetChatAdministratorsParams) ([]models.ChatMember, error) {
	var result []models.ChatMember
	err := b.rawRequest(ctx, "getChatAdministrators", params, &result)
	return result, err
}

// GetChatMemberCount https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(ctx context.Context, params *GetChatMemberCountParams) (int, error) {
	var result int
	err := b.rawRequest(ctx, "getChatMemberCount", params, &result)
	return result, err
}

// GetChatMember https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(ctx context.Context, params *GetChatMemberParams) (*models.ChatMember, error) {
	result := &models.ChatMember{}
	err := b.rawRequest(ctx, "getChatMember", params, &result)
	return result, err
}

// SetChatStickerSet https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(ctx context.Context, params *SetChatStickerSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatStickerSet", params, &result)
	return result, err
}

// GetForumTopicIconStickers https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers(ctx context.Context) ([]*models.Sticker, error) {
	var result []*models.Sticker
	err := b.rawRequest(ctx, "getForumTopicIconStickers", nil, &result)
	return result, err
}

// CreateForumTopic https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(ctx context.Context, params *CreateForumTopicParams) (*models.ForumTopic, error) {
	result := &models.ForumTopic{}
	err := b.rawRequest(ctx, "createForumTopic", params, &result)
	return result, err
}

// EditForumTopic https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(ctx context.Context, params *EditForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "editForumTopic", params, &result)
	return result, err
}

// CloseForumTopic https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(ctx context.Context, params *CloseForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "closeForumTopic", params, &result)
	return result, err
}

// ReopenForumTopic https://core.telegram.org/bots/api#reopenforumtopic
func (b *Bot) ReopenForumTopic(ctx context.Context, params *ReopenForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "reopenForumTopic", params, &result)
	return result, err
}

// UnpinAllForumTopicMessages https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) UnpinAllForumTopicMessages(ctx context.Context, params *UnpinAllForumTopicMessagesParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unpinAllForumTopicMessages", params, &result)
	return result, err
}

// EditGeneralForumTopic https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(ctx context.Context, params *EditGeneralForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "editGeneralForumTopic", params, &result)
	return result, err
}

// CloseGeneralForumTopic https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(ctx context.Context, params *CloseGeneralForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "closeGeneralForumTopic", params, &result)
	return result, err
}

// ReopenGeneralForumTopic https://core.telegram.org/bots/api#reopengeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(ctx context.Context, params *ReopenGeneralForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "reopenGeneralForumTopic", params, &result)
	return result, err
}

// HideGeneralForumTopic https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(ctx context.Context, params *HideGeneralForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "hideGeneralForumTopic", params, &result)
	return result, err
}

// UnhideGeneralForumTopic https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(ctx context.Context, params *UnhideGeneralForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unhideGeneralForumTopic", params, &result)
	return result, err
}

// UnpinAllGeneralForumTopicMessages https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, params *UnpinAllGeneralForumTopicMessagesParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "unpinAllGeneralForumTopicMessages", params, &result)
	return result, err
}

// DeleteForumTopic https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) DeleteForumTopic(ctx context.Context, params *DeleteForumTopicParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteForumTopic", params, &result)
	return result, err
}

// DeleteChatStickerSet https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(ctx context.Context, params *DeleteChatStickerSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteChatStickerSet", params, &result)
	return result, err
}

// AnswerCallbackQuery https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(ctx context.Context, params *AnswerCallbackQueryParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "answerCallbackQuery", params, &result)
	return result, err
}

// GetUserChatBoosts https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(ctx context.Context, params *GetUserChatBoostsParams) (*models.UserChatBoosts, error) {
	result := &models.UserChatBoosts{}
	err := b.rawRequest(ctx, "getUserChatBoosts", params, &result)
	return result, err
}

// GetBusinessConnection https://core.telegram.org/bots/api#getbusinessconnection
func (b *Bot) GetBusinessConnection(ctx context.Context, params *GetBusinessConnectionParams) (*models.BusinessConnection, error) {
	result := &models.BusinessConnection{}
	err := b.rawRequest(ctx, "getBusinessConnection", params, &result)
	return result, err
}

// SetMyCommands https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(ctx context.Context, params *SetMyCommandsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMyCommands", params, &result)
	return result, err
}

// DeleteMyCommands https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(ctx context.Context, params *DeleteMyCommandsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteMyCommands", params, &result)
	return result, err
}

// GetMyCommands https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(ctx context.Context, params *GetMyCommandsParams) ([]models.BotCommand, error) {
	var result []models.BotCommand
	err := b.rawRequest(ctx, "getMyCommands", params, &result)
	return result, err
}

// SetMyName https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(ctx context.Context, params *SetMyNameParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMyName", params, &result)
	return result, err
}

// GetMyName https://core.telegram.org/bots/api#getmyname
func (b *Bot) GetMyName(ctx context.Context, params *GetMyNameParams) (models.BotName, error) {
	var result models.BotName
	err := b.rawRequest(ctx, "getMyName", params, &result)
	return result, err
}

// SetMyDescription https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(ctx context.Context, params *SetMyDescriptionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMyDescription", params, &result)
	return result, err
}

// GetMyDescription https://core.telegram.org/bots/api#getmydescription
func (b *Bot) GetMyDescription(ctx context.Context, params *GetMyDescriptionParams) (models.BotDescription, error) {
	var result models.BotDescription
	err := b.rawRequest(ctx, "getMyDescription", params, &result)
	return result, err
}

// SetMyShortDescription https://core.telegram.org/bots/api#setmyshortdescription
func (b *Bot) SetMyShortDescription(ctx context.Context, params *SetMyShortDescriptionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMyShortDescription", params, &result)
	return result, err
}

// GetMyShortDescription https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(ctx context.Context, params *GetMyShortDescriptionParams) (models.BotShortDescription, error) {
	var result models.BotShortDescription
	err := b.rawRequest(ctx, "getMyShortDescription", params, &result)
	return result, err
}

// SetChatMenuButton https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(ctx context.Context, params *SetChatMenuButtonParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setChatMenuButton", params, &result)
	return result, err
}

// GetChatMenuButton https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(ctx context.Context, params *GetChatMenuButtonParams) (models.MenuButton, error) {
	var result models.MenuButton
	err := b.rawRequest(ctx, "getChatMenuButton", params, &result)
	return result, err
}

// SetMyDefaultAdministratorRights https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(ctx context.Context, params *SetMyDefaultAdministratorRightsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setMyDefaultAdministratorRights", params, &result)
	return result, err
}

// GetMyDefaultAdministratorRights https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(ctx context.Context, params *GetMyDefaultAdministratorRightsParams) (*models.ChatAdministratorRights, error) {
	result := &models.ChatAdministratorRights{}
	err := b.rawRequest(ctx, "setMyDefaultAdministratorRights", params, &result)
	return result, err
}

// EditMessageText https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(ctx context.Context, params *EditMessageTextParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "editMessageText", params, result)
	return result, err
}

// EditMessageCaption https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(ctx context.Context, params *EditMessageCaptionParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "editMessageCaption", params, result)
	return result, err
}

// EditMessageMedia https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(ctx context.Context, params *EditMessageMediaParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "editMessageMedia", params, result)
	return result, err
}

// EditMessageReplyMarkup https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(ctx context.Context, params *EditMessageReplyMarkupParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "editMessageReplyMarkup", params, result)
	return result, err
}

// StopPoll https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(ctx context.Context, params *StopPollParams) (*models.Poll, error) {
	result := &models.Poll{}
	err := b.rawRequest(ctx, "stopPoll", params, &result)
	return result, err
}

// DeleteMessage https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(ctx context.Context, params *DeleteMessageParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteMessage", params, &result)
	return result, err
}

// DeleteMessages https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(ctx context.Context, params *DeleteMessagesParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteMessages", params, &result)
	return result, err
}

// SendSticker https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(ctx context.Context, params *SendStickerParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendSticker", params, &result)
	return result, err
}

// GetStickerSet https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(ctx context.Context, params *GetStickerSetParams) (*models.StickerSet, error) {
	result := &models.StickerSet{}
	err := b.rawRequest(ctx, "getStickerSet", params, &result)
	return result, err
}

// GetCustomEmojiStickers https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(ctx context.Context, params *GetCustomEmojiStickersParams) ([]*models.Sticker, error) {
	var result []*models.Sticker
	err := b.rawRequest(ctx, "getCustomEmojiStickers", params, &result)
	return result, err
}

// UploadStickerFile https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(ctx context.Context, params *UploadStickerFileParams) (*models.File, error) {
	result := &models.File{}
	err := b.rawRequest(ctx, "uploadStickerFile", params, &result)
	return result, err
}

// CreateNewStickerSet https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(ctx context.Context, params *CreateNewStickerSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "createNewStickerSet", params, &result)
	return result, err
}

// AddStickerToSet https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(ctx context.Context, params *AddStickerToSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "addStickerToSet", params, &result)
	return result, err
}

// SetStickerPositionInSet https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(ctx context.Context, params *SetStickerPositionInSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerPositionInSet", params, &result)
	return result, err
}

// DeleteStickerFromSet https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(ctx context.Context, params *DeleteStickerFromSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteStickerFromSet", params, &result)
	return result, err
}

// ReplaceStickerInSet https://core.telegram.org/bots/api#replacestickerinset
func (b *Bot) ReplaceStickerInSet(ctx context.Context, params *ReplaceStickerInSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "replaceStickerInSet", params, &result)
	return result, err
}

// SetStickerEmojiList https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(ctx context.Context, params *SetStickerEmojiListParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerEmojiList", params, &result)
	return result, err
}

// SetStickerKeywords https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(ctx context.Context, params *SetStickerKeywordsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerKeywords", params, &result)
	return result, err
}

// SetStickerMaskPosition https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(ctx context.Context, params *SetStickerMaskPositionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerMaskPosition", params, &result)
	return result, err
}

// SetStickerSetTitle https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerSetTitle(ctx context.Context, params *SetStickerSetTitleParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerSetTitle", params, &result)
	return result, err
}

// SetStickerSetThumbnail https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(ctx context.Context, params *SetStickerSetThumbnailParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setStickerSetThumbnail", params, &result)
	return result, err
}

// SetCustomEmojiStickerSetThumbnail https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, params *SetCustomEmojiStickerSetThumbnailParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setCustomEmojiStickerSetThumbnail", params, &result)
	return result, err
}

// DeleteStickerSet https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(ctx context.Context, params *DeleteStickerSetParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteStickerSet", params, &result)
	return result, err
}

// AnswerInlineQuery https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(ctx context.Context, params *AnswerInlineQueryParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "answerInlineQuery", params, &result)
	return result, err
}

// AnswerWebAppQuery https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(ctx context.Context, params *AnswerWebAppQueryParams) (*models.SentWebAppMessage, error) {
	result := &models.SentWebAppMessage{}
	err := b.rawRequest(ctx, "answerWebAppQuery", params, result)
	return result, err
}

// SavePreparedInlineMessage https://core.telegram.org/bots/api#savepreparedinlinemessage
func (b *Bot) SavePreparedInlineMessage(ctx context.Context, params *SavePreparedInlineMessageParams) (*models.PreparedInlineMessage, error) {
	result := &models.PreparedInlineMessage{}
	err := b.rawRequest(ctx, "savePreparedInlineMessage", params, result)
	return result, err
}

// SendInvoice https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(ctx context.Context, params *SendInvoiceParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendInvoice", params, result)
	return result, err
}

// CreateInvoiceLink https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(ctx context.Context, params *CreateInvoiceLinkParams) (string, error) {
	var result string
	err := b.rawRequest(ctx, "createInvoiceLink", params, &result)
	return result, err
}

// AnswerShippingQuery https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(ctx context.Context, params *AnswerShippingQueryParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "answerShippingQuery", params, &result)
	return result, err
}

// AnswerPreCheckoutQuery https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, params *AnswerPreCheckoutQueryParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "answerPreCheckoutQuery", params, &result)
	return result, err
}

// GetStarTransactions https://core.telegram.org/bots/api#getstartransactions
func (b *Bot) GetStarTransactions(ctx context.Context, params *GetStarTransactionsParams) (*models.StarTransactions, error) {
	result := models.StarTransactions{}
	err := b.rawRequest(ctx, "getStarTransactions", params, &result)
	return &result, err
}

// RefundStarPayment https://core.telegram.org/bots/api#refundstarpayment
func (b *Bot) RefundStarPayment(ctx context.Context, params *RefundStarPaymentParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "refundStarPayment", params, &result)
	return result, err
}

// EditUserStarSubscription https://core.telegram.org/bots/api#edituserstarsubscription
func (b *Bot) EditUserStarSubscription(ctx context.Context, params *EditUserStarSubscriptionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "editUserStarSubscription", params, &result)
	return result, err
}

// SetPassportDataErrors https://core.telegram.org/bots/api#setpassportdataerrors
func (b *Bot) SetPassportDataErrors(ctx context.Context, params *SetPassportDataErrorsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setPassportDataErrors", params, &result)
	return result, err
}

// SendGame https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(ctx context.Context, params *SendGameParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "sendGame", params, result)
	return result, err
}

// SetGameScore https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(ctx context.Context, params *SetGameScoreParams) (*models.Message, error) {
	result := &models.Message{}
	err := b.rawRequest(ctx, "setGameScore", params, result)
	return result, err
}

// GetGameHighScores https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(ctx context.Context, params *GetGameHighScoresParams) ([]*models.GameHighScore, error) {
	var result []*models.GameHighScore
	err := b.rawRequest(ctx, "getGameHighScores", params, &result)
	return result, err
}

// GetAvailableGifts https://core.telegram.org/bots/api#getavailablegifts
func (b *Bot) GetAvailableGifts(ctx context.Context) (*models.Gifts, error) {
	result := &models.Gifts{}
	err := b.rawRequest(ctx, "getAvailableGifts", nil, result)
	return result, err
}

// SendGift https://core.telegram.org/bots/api#sendgift
func (b *Bot) SendGift(ctx context.Context, params *SendGiftParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "sendGift", params, &result)
	return result, err
}

// VerifyUser https://core.telegram.org/bots/api#verifyuser
func (b *Bot) VerifyUser(ctx context.Context, params *VerifyUserParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "verifyUser", params, &result)
	return result, err
}

// VerifyChat https://core.telegram.org/bots/api#verifychat
func (b *Bot) VerifyChat(ctx context.Context, params *VerifyChatParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "verifyChat", params, &result)
	return result, err
}

// RemoveUserVerification https://core.telegram.org/bots/api#removeuserverification
func (b *Bot) RemoveUserVerification(ctx context.Context, params *RemoveUserVerificationParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "removeUserVerification", params, &result)
	return result, err
}

// RemoveChatVerification https://core.telegram.org/bots/api#removechatverification
func (b *Bot) RemoveChatVerification(ctx context.Context, params *RemoveChatVerificationParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "removeChatVerification", params, &result)
	return result, err
}

// ReadBusinessMessage https://core.telegram.org/bots/api#readbusinessmessage
func (b *Bot) ReadBusinessMessage(ctx context.Context, params *ReadBusinessMessageParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "readBusinessMessage", params, &result)
	return result, err
}

// DeleteBusinessMessages https://core.telegram.org/bots/api#deletebusinessmessages
func (b *Bot) DeleteBusinessMessages(ctx context.Context, params *DeleteBusinessMessagesParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteBusinessMessages", params, &result)
	return result, err
}

// SetBusinessAccountName https://core.telegram.org/bots/api#setbusinessaccountname
func (b *Bot) SetBusinessAccountName(ctx context.Context, params *SetBusinessAccountNameParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setBusinessAccountName", params, &result)
	return result, err
}

// SetBusinessAccountUsername https://core.telegram.org/bots/api#setbusinessaccountusername
func (b *Bot) SetBusinessAccountUsername(ctx context.Context, params *SetBusinessAccountUsernameParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setBusinessAccountUsername", params, &result)
	return result, err
}

// SetBusinessAccountBio https://core.telegram.org/bots/api#setbusinessaccountbio
func (b *Bot) SetBusinessAccountBio(ctx context.Context, params *SetBusinessAccountBioParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setBusinessAccountBio", params, &result)
	return result, err
}

// SetBusinessAccountProfilePhoto https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (b *Bot) SetBusinessAccountProfilePhoto(ctx context.Context, params *SetBusinessAccountProfilePhotoParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "SetBusinessAccountProfilePhoto", params, &result)
	return result, err
}

// RemoveBusinessAccountProfilePhoto https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (b *Bot) RemoveBusinessAccountProfilePhoto(ctx context.Context, params *RemoveBusinessAccountProfilePhotoParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "removeBusinessAccountProfilePhoto", params, &result)
	return result, err
}

// SetBusinessAccountGiftSettings https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (b *Bot) SetBusinessAccountGiftSettings(ctx context.Context, params *SetBusinessAccountGiftSettingsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "setBusinessAccountGiftSettings", params, &result)
	return result, err
}

// GetBusinessAccountStarBalance https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (b *Bot) GetBusinessAccountStarBalance(ctx context.Context, params *GetBusinessAccountStarBalanceParams) (*models.StarAmount, error) {
	result := &models.StarAmount{}
	err := b.rawRequest(ctx, "getBusinessAccountStarBalance", params, &result)
	return result, err
}

// TransferBusinessAccountStars https://core.telegram.org/bots/api#transferbusinessaccountstars
func (b *Bot) TransferBusinessAccountStars(ctx context.Context, params *TransferBusinessAccountStarsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "transferBusinessAccountStars", params, &result)
	return result, err
}

// GetBusinessAccountGifts https://core.telegram.org/bots/api#getbusinessaccountgifts
func (b *Bot) GetBusinessAccountGifts(ctx context.Context, params *GetBusinessAccountGiftsParams) (*models.OwnedGifts, error) {
	result := &models.OwnedGifts{}
	err := b.rawRequest(ctx, "getBusinessAccountGifts", params, &result)
	return result, err
}

// ConvertGiftToStars https://core.telegram.org/bots/api#convertgifttostars
func (b *Bot) ConvertGiftToStars(ctx context.Context, params *ConvertGiftToStarsParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "convertGiftToStars", params, &result)
	return result, err
}

// UpgradeGift https://core.telegram.org/bots/api#upgradegift
func (b *Bot) UpgradeGift(ctx context.Context, params *UpgradeGiftParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "upgradeGift", params, &result)
	return result, err
}

// TransferGift https://core.telegram.org/bots/api#transfergift
func (b *Bot) TransferGift(ctx context.Context, params *TransferGiftParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "transferGift", params, &result)
	return result, err
}

// PostStory https://core.telegram.org/bots/api#poststory
func (b *Bot) PostStory(ctx context.Context, params *PostStoryParams) (*models.Story, error) {
	result := &models.Story{}
	err := b.rawRequest(ctx, "postStory", params, &result)
	return result, err
}

// EditStory https://core.telegram.org/bots/api#editstory
func (b *Bot) EditStory(ctx context.Context, params *EditStoryParams) (*models.Story, error) {
	result := &models.Story{}
	err := b.rawRequest(ctx, "editStory", params, &result)
	return result, err
}

// DeleteStory https://core.telegram.org/bots/api#deletestory
func (b *Bot) DeleteStory(ctx context.Context, params *DeleteStoryParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "deleteStory", params, &result)
	return result, err
}

// GiftPremiumSubscription https://core.telegram.org/bots/api#giftpremiumsubscription
func (b *Bot) GiftPremiumSubscription(ctx context.Context, params *GiftPremiumSubscriptionParams) (bool, error) {
	var result bool
	err := b.rawRequest(ctx, "giftPremiumSubscription", params, &result)
	return result, err
}
