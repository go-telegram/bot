package keyboards

import "github.com/go-telegram/bot/models"

type InlineKeyboard struct {
	rows [][]models.InlineKeyboardButton
}

func NewInlineKeyboard() *InlineKeyboard {
	return &InlineKeyboard{}
}

func (b *InlineKeyboard) Row(btns ...models.InlineKeyboardButton) *InlineKeyboard {
	b.rows = append(b.rows, btns)
	return b
}

func (b *InlineKeyboard) Build() models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: b.rows,
	}
}

func InlineKeyboardButton(text, data string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

func InlineKeyboardURL(text, url string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

func InlineKeyboardWebApp(text string, webApp models.WebAppInfo) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:   text,
		WebApp: &webApp,
	}
}

func InlineKeyboardLoginURL(text string, loginURL models.LoginURL) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:     text,
		LoginURL: &loginURL,
	}
}

func InlineKeyboardSwitchInlineQuery(text, query string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

func InlineKeyboardSwitchInlineQueryCurrentChat(text, query string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:                         text,
		SwitchInlineQueryCurrentChat: query,
	}
}

func InlineKeyboardSwitchInlineQueryChosenChat(text string, chosenChat models.SwitchInlineQueryChosenChat) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:                        text,
		SwitchInlineQueryChosenChat: &chosenChat,
	}
}

func InlineKeyboardCopyText(text string, copyText models.CopyTextButton) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:     text,
		CopyText: copyText,
	}
}

func InlineKeyboardCallbackGame(text string, callbackGame models.CallbackGame) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:         text,
		CallbackGame: &callbackGame,
	}
}

func InlineKeyboardPay(text string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}
