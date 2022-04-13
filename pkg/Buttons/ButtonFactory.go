package Buttons

import (
	tb "gopkg.in/telebot.v3"
	"strconv"
)

type Icon string
type Command string

type Factory struct {
	ToyButtonTemplates map[string]string
}

func (factory *Factory) GenerateToyCommandButtons(toyButtons map[string]*tb.Btn, unificator int, commands []string) {
	for _, command := range commands {
		toyButtons[command] = &tb.Btn{Unique: command + strconv.Itoa(unificator), Text: factory.ToyButtonTemplates[command]}
	}
}
