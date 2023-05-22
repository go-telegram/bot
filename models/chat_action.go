package models

type ChatAction string

const (
	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto                = "upload_photo"
	ChatActionRecordVideo                = "record_video"
	ChatActionUploadVideo                = "upload_video"
	ChatActionRecordVoice                = "record_voice"
	ChatActionUploadVoice                = "upload_voice"
	ChatActionUploadDocument             = "upload_document"
	ChatActionChooseSticker              = "choose_sticker"
	ChatActionFindLocation               = "find_location"
	ChatActionRecordVideoNote            = "record_video_note"
	ChatActionUploadVideoNote            = "upload_video_note"
)
