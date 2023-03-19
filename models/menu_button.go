package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MenuButtonType int

const (
	MenuButtonTypeCommands MenuButtonType = iota
	MenuButtonTypeWebApp
	MenuButtonTypeDefault
)

type InputMenuButton interface {
	menuButtonTag()
}

// MenuButton https://core.telegram.org/bots/api#menubutton
type MenuButton struct {
	Type MenuButtonType

	Commands *MenuButtonCommands
	WebApp   *MenuButtonWebApp
	Default  *MenuButtonDefault
}

func (c *MenuButton) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, []byte(`"type":"commands"`)) {
		c.Type = MenuButtonTypeCommands
		c.Commands = &MenuButtonCommands{}
		return json.Unmarshal(data, c.Commands)
	}
	if bytes.Contains(data, []byte(`"type":"web_app"`)) {
		c.Type = MenuButtonTypeWebApp
		c.WebApp = &MenuButtonWebApp{}
		return json.Unmarshal(data, c.WebApp)
	}
	if bytes.Contains(data, []byte(`"type":"default"`)) {
		c.Type = MenuButtonTypeDefault
		c.Default = &MenuButtonDefault{}
		return json.Unmarshal(data, c.Default)
	}

	return fmt.Errorf("unsupported MenuButton type")
}

// MenuButtonCommands https://core.telegram.org/bots/api#menubuttoncommands
type MenuButtonCommands struct {
	Type string `json:"type" rules:"required,equals:commands"`
}

func (MenuButtonCommands) menuButtonTag() {}

// MenuButtonWebApp https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	Type   string     `json:"type" rules:"required,equals:web_app"`
	Text   string     `json:"text" rules:"required"`
	WebApp WebAppInfo `json:"web_app" rules:"required"`
}

func (MenuButtonWebApp) menuButtonTag() {}

// MenuButtonDefault https://core.telegram.org/bots/api#menubuttondefault
type MenuButtonDefault struct {
	Type string `json:"type" rules:"required,equals:default"`
}

func (MenuButtonDefault) menuButtonTag() {}
