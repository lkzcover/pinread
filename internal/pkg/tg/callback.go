package tg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/pkg/tg/command"
	"github.com/lkzcover/tapi"
	"strings"
)

func callbackRouter(ctx context.Context, db *pgxpool.Pool, tApi *tapi.Engine, msg tapi.Message) error {
	commandList := strings.Split(*msg.CallbackQuery.Data, "&")
	switch commandList[0] {
	case command.Read, command.Unexciting:
		return actionRead(ctx, db, tApi, msg, commandList)
	default:
		// TODO alert for unknown command @p.novokshonov
		return nil
	}
}
