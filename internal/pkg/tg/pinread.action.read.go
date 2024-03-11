package tg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/pkg/repository"
	"github.com/lkzcover/tapi"
	"strconv"
)

func actionRead(ctx context.Context, db *pgxpool.Pool, tApi *tapi.Engine, msg tapi.Message, command []string) error {
	msgID, err := strconv.ParseUint(command[1], 10, 64)
	if err != nil {
		return err // TODO details error @lkzcover
	}

	err = repository.NewDBPinsRepoRepo(db).DeactivatePin(ctx, msg.CallbackQuery.From.ID, msgID)
	if err != nil {
		return err
	}

	// TODO not good way @lkzcover
	_, err = tApi.Reply(&tapi.Message{
		Message: tapi.MessageStruct{
			MessageID: msgID,
			Chat: struct {
				ID int64 `json:"id"`
			}(struct{ ID int64 }{
				ID: msg.CallbackQuery.Message.Chat.ID,
			}),
		},
	}, tapi.MsgParams{Text: "success"})
	if err != nil {
		return err
	}

	return nil
}
