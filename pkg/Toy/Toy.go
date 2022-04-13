package Toy

import (
	tb "gopkg.in/telebot.v3"
)

type Toy struct {
	Name              string
	ID                int
	AvailableCommands []string
	Buttons           map[string]*tb.Btn
	PublishTopic      string
	SubscribeTopic    string
	KeyboardName      string
}
