package Buttons

import (
	tb "gopkg.in/telebot.v3"
	"strconv"
)

type Icon string
type Command string

type Factory struct {
	ToyButtonTemplates map[Command]Icon
}

func (factory *Factory) GenerateToyCommandButtons(toyButtons *[]tb.Btn, unificator int, commands []string) {
	for _, command := range commands {
		func() {
			*toyButtons = append(*toyButtons, tb.Btn{Unique: command + strconv.Itoa(unificator), Text: string(factory.ToyButtonTemplates[Command(command)])})
		}()
	}
}
