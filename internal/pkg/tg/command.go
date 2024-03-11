package tg

import (
	"github.com/lkzcover/pinread/internal/pkg/tg/command"
	"github.com/lkzcover/pinread/internal/pkg/tg/tg_interface"
	"github.com/lkzcover/tapi"
)

func commandRouter(tApi *tapi.Engine, msg tapi.Message) error {
	switch msg.Message.Text {
	case command.Start:
		{
			tApi.SendMessage(msg.Message.Chat.ID, tapi.MsgParams{Text: "hello"}, tg_interface.ButtonMain())
			return nil
		}
	default:
		// TODO alert for unknown command @p.novokshonov
		return nil
	}
}
