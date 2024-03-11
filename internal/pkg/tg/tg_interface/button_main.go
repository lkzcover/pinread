package tg_interface

import "github.com/lkzcover/tapi"

func ButtonMain() tapi.ReplyKeyboardMarkup {
	return tapi.ReplyKeyboardMarkup{
		Keyboard: [][]tapi.KeyboardButton{
			{
				{Text: "🎲 get random pin"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
		Selective:       true,
	}
}
