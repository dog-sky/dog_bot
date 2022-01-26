package db

import (
	"context"
	"time"

	"github.com/dog-sky/dog_bot/internal/service/db/models"
	desc "github.com/dog-sky/dog_bot/pkg/dog/api"
	"google.golang.org/protobuf/types/known/timestamppb"

	sq "github.com/Masterminds/squirrel"
)

const (
	interactionTable = "interaction"

	interactionTableStatusColumn = "interaction_status"
	cretedAtColumn               = "created_at"
)

func (d *db) setStatusQuery(ctx context.Context, status string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Insert(interactionTable).
		Columns(interactionTableStatusColumn).
		Values(status).
		ToSql()
	if err != nil {
		return err
	}

	c, err := d.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer c.Release()

	_, err = c.Exec(ctx, sql, args...)

	return err
}

func (d *db) statusListQuery(ctx context.Context, filter *models.Filter) ([]*desc.StatusListReply_Action, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	res := []*desc.StatusListReply_Action{}

	builder := psql.
		Select(interactionTableStatusColumn, cretedAtColumn).
		From(interactionTable)

	if len(filter.Actions) != 0 {
		builder = builder.Where(sq.Eq{interactionTableStatusColumn: filter.Actions})
	}
	if filter.Date.From != nil {
		builder = builder.Where(sq.GtOrEq{cretedAtColumn: filter.Date.From})
	}
	if filter.Date.To != nil {
		builder = builder.Where(sq.LtOrEq{cretedAtColumn: filter.Date.To})
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return res, err
	}

	c, err := d.pool.Acquire(ctx)
	if err != nil {
		return res, err
	}

	defer c.Release()

	rows, err := c.Query(ctx, sql, args...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		var action string
		var date time.Time

		err = rows.Scan(&action, &date)
		if err != nil {
			return res, err
		}

		res = append(res, &desc.StatusListReply_Action{
			Action: desc.DogAction(desc.DogAction_value[action]),
			Date:   timestamppb.New(date),
		})
	}

	return res, nil
}
