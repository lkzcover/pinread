package repository

import (
	"context"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/generated/pinread/public/model"
	"github.com/lkzcover/pinread/internal/generated/pinread/public/table"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

const tablePins = "pins."

type DBPinsRepo struct {
	db *pgxpool.Pool
}

func NewDBPinsRepoRepo(db *pgxpool.Pool) DBPinsRepo {
	return DBPinsRepo{db: db}
}

func (obj DBPinsRepo) AddNewPin(ctx context.Context, pin model.Pins) error {
	query, args := table.Pins.INSERT(table.Pins.AllColumns).MODEL(pin).Sql()

	_, err := obj.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (obj DBPinsRepo) DeactivatePin(ctx context.Context, userID int64, msgID uint64) error {
	query, args := table.Pins.UPDATE(table.Pins.IsActive).
		SET(postgres.Bool(false)).
		WHERE(
			table.Pins.UserID.EQ(postgres.Int64(userID)).
				AND(table.Pins.MsgID.EQ(postgres.Uint64(msgID)))).Sql()

	_, err := obj.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (obj DBPinsRepo) GetRandomUnread(ctx context.Context, userID int64) ([]model.Pins, error) {
	query, args := table.Pins.SELECT(table.Pins.AllColumns).
		WHERE(
			table.Pins.UserID.EQ(postgres.Int64(userID)).
				AND(table.Pins.IsActive.IS_TRUE())).Sql()

	pins, err := pgxscan.All(ctx, obj.db, scan.StructMapper[model.Pins](scan.WithStructTagPrefix("pins.")), query, args...)

	return pins, err
}
