package tg

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/lkzcover/pinread/internal/pkg/tg/tg_interface"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/pkg/repository"
	"github.com/lkzcover/tapi"
)

func getRandomPin(ctx context.Context, db *pgxpool.Pool, tApi *tapi.Engine, msg tapi.Message) {
	pins, err := repository.NewDBPinsRepoRepo(db).GetRandomUnread(ctx, msg.Message.Chat.ID)
	if err != nil {
		log.Printf("get active pins error: %s", err)
		return
	}

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	pinN := rnd.Intn(len(pins))
	pinN = pinN - 1

	// TODO not good way @lkzcover
	_, err = tApi.Reply(&tapi.Message{
		Message: tapi.MessageStruct{
			MessageID: uint64(pins[pinN].MsgID),
			Chat: struct {
				ID int64 `json:"id"`
			}(struct{ ID int64 }{
				ID: pins[pinN].UserID,
			}),
		},
	}, tapi.MsgParams{Text: "let's read"}, tg_interface.InlineDefaultPin(uint64(pins[pinN].MsgID)))
	if err != nil {
		log.Printf("send random pin error: %s", err)
		return
	}

	return
}
