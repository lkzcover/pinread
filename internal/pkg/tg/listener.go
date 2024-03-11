package tg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/pkg/env"
	"github.com/lkzcover/tapi"
	"log"
	"strings"
	"time"
)

// StartListener get new message from telegram_api
func StartListener(db *pgxpool.Pool) {
	go func() {
		tgConn := tapi.Init(env.Setup.TelegramToken)
		ctx := context.Background()

		for {
			msgs, err := tgConn.GetUpdates()
			if err != nil {
				log.Printf("ERROR get message from telegram error: %s", err)
				time.Sleep(20 * time.Second)
				continue
			}

			for _, msg := range msgs {
				if msg.CallbackQuery != nil {
					err = callbackRouter(ctx, db, &tgConn, msg)
					if err != nil {
						log.Printf("ERROR callback engine error: %s", err)
					}

					continue
				}

				if strings.HasPrefix(msg.Message.Text, "/") {
					err = commandRouter(&tgConn, msg)
					if err != nil {
						log.Printf("ERROR command engine error: %s", err)
					}

					continue
				}

				if strings.HasPrefix(msg.Message.Text, "🎲") {
					go getRandomPin(ctx, db, &tgConn, msg)

					continue
				}

				// ADD new pinread
				if strings.HasPrefix(msg.Message.Text, "http") {
					go addNewPinRead(db, &tgConn, msg)
				}

				//TODO add regexp for correct router @p.novokshonov

			}

			time.Sleep(10 * time.Second) // TODO add dynamic timeout ??? @lkzcover
		}
	}()
}
