package models

type ParseMode string

const (
	ParseModeMarkdownV1 ParseMode = "Markdown"
	ParseModeMarkdown   ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)
