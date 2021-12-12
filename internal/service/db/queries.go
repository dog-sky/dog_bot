package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

const (
	interactionTable = "interaction"

	interactionTableStatus = "interaction_status"
)

func (d *db) setStatusQuery(ctx context.Context, status string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Insert(interactionTable).
		Columns(interactionTableStatus).
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
