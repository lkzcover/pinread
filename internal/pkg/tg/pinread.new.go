package tg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/generated/pinread/public/model"
	"github.com/lkzcover/pinread/internal/pkg/repository"
	"github.com/lkzcover/pinread/internal/pkg/tg/tg_interface"
	"github.com/lkzcover/tapi"
	"log"
)

func addNewPinRead(db *pgxpool.Pool, tApi *tapi.Engine, msg tapi.Message) {
	ctx := context.Background()
	err := repository.NewDBPinsRepoRepo(db).AddNewPin(ctx, model.Pins{
		UserID:   msg.Message.From.ID,
		MsgID:    int64(msg.Message.MessageID),
		IsActive: true,
	})
	if err != nil {
		log.Printf("add new pin error: %s", err)
		return
	}

	_, err = tApi.Reply(&msg, tapi.MsgParams{
		Text: "add to main list",
	}, tg_interface.InlineDefaultPin(msg.Message.MessageID))
	if err != nil {
		log.Printf("reply msg error: %s", err)
	}

	return
}
