package Keyboards

import tb "gopkg.in/telebot.v3"

type KeyboardFactory struct {
	MenuButtonTemplates map[string]string
}

func (factory *KeyboardFactory) GenerateToyCommandsKeyboard(toyButtons map[string]*tb.Btn) (toyKeyboard tb.ReplyMarkup) {
	var buttonsSlice = make([]tb.Btn, len(toyButtons))

	i := 0
	for name, _ := range toyButtons {
		buttonsSlice[i] = *toyButtons[name]
		i++
	}

	toyKeyboard = tb.ReplyMarkup{ResizeKeyboard: true}
	toyKeyboard.Inline(toyKeyboard.Row(buttonsSlice...))
	return toyKeyboard
}
