package models

import (
	"encoding/json"
	"fmt"
)

type MenuButtonType string

const (
	MenuButtonTypeCommands MenuButtonType = "commands"
	MenuButtonTypeWebApp   MenuButtonType = "web_app"
	MenuButtonTypeDefault  MenuButtonType = "default"
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
	v := struct {
		Type MenuButtonType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case MenuButtonTypeCommands:
		c.Type = MenuButtonTypeCommands
		c.Commands = &MenuButtonCommands{}
		return json.Unmarshal(data, c.Commands)
	case MenuButtonTypeWebApp:
		c.Type = MenuButtonTypeWebApp
		c.WebApp = &MenuButtonWebApp{}
		return json.Unmarshal(data, c.WebApp)
	case MenuButtonTypeDefault:
		c.Type = MenuButtonTypeDefault
		c.Default = &MenuButtonDefault{}
		return json.Unmarshal(data, c.Default)
	}

	return fmt.Errorf("unsupported MenuButton type")
}

func (c *MenuButton) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case MenuButtonTypeCommands:
		c.Commands.Type = MenuButtonTypeCommands
		return json.Marshal(c.Commands)
	case MenuButtonTypeWebApp:
		c.WebApp.Type = MenuButtonTypeWebApp
		return json.Marshal(c.WebApp)
	case MenuButtonTypeDefault:
		c.Default.Type = MenuButtonTypeDefault
		return json.Marshal(c.Default)
	}

	return nil, fmt.Errorf("unsupported MenuButton type")
}

// MenuButtonCommands https://core.telegram.org/bots/api#menubuttoncommands
type MenuButtonCommands struct {
	Type MenuButtonType `json:"type" rules:"required,equals:commands"`
}

func (MenuButtonCommands) menuButtonTag() {}

// MenuButtonWebApp https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	Type   MenuButtonType `json:"type" rules:"required,equals:web_app"`
	Text   string         `json:"text" rules:"required"`
	WebApp WebAppInfo     `json:"web_app" rules:"required"`
}

func (MenuButtonWebApp) menuButtonTag() {}

// MenuButtonDefault https://core.telegram.org/bots/api#menubuttondefault
type MenuButtonDefault struct {
	Type MenuButtonType `json:"type" rules:"required,equals:default"`
}

func (MenuButtonDefault) menuButtonTag() {}
