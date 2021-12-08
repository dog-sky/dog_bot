package db

//go:generate mkdir -p mocks
//go:generate rm -rf ./mocks/*_minimock.go
//go:generate minimock -i DB -o ./mocks/ -s "_minimock.go"

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	SetStatus(status string) error

	ShutDown()
}

type db struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, dbURL string) (DB, error) {
	conn, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	return &db{
		pool: conn,
	}, nil
}

func (d *db) ShutDown() {
	d.pool.Close()
}

func (d *db) SetStatus(status string) error {
	// TODO сделать запрос и положить данные в базу
	return nil
}
