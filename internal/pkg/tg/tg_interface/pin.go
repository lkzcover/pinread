package tg_interface

import (
	"fmt"
	"github.com/lkzcover/pinread/internal/pkg/tg/command"
	"github.com/lkzcover/tapi"
)

func InlineDefaultPin(msgID uint64) tapi.InlineKeyboardMarkup {
	var buttonLine []tapi.InlineKeyboardButton

	read := fmt.Sprintf("%s&%d", command.Read, msgID)
	unex := fmt.Sprintf("%s&%d", command.Unexciting, msgID)
	later := fmt.Sprintf("%s&%d", command.Later, msgID)

	buttonLine = append(buttonLine, tapi.InlineKeyboardButton{Text: "✅ Read", CallbackData: &read})
	buttonLine = append(buttonLine, tapi.InlineKeyboardButton{Text: "🕓 Later", CallbackData: &later})
	buttonLine = append(buttonLine, tapi.InlineKeyboardButton{Text: "❌ Unexciting", CallbackData: &unex})

	return tapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tapi.InlineKeyboardButton{buttonLine},
	}
}
