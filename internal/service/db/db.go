package db

//go:generate mkdir -p mocks
//go:generate rm -rf ./mocks/*_minimock.go
//go:generate minimock -i DB -o ./mocks/ -s "_minimock.go"

import (
	"context"

	"github.com/dog-sky/dog_bot/internal/service/db/models"
	desc "github.com/dog-sky/dog_bot/pkg/dog/api"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	SetStatus(ctx context.Context, status string) error
	StatusList(ctx context.Context, filter *models.Filter) ([]*desc.StatusListReply_Action, error)

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

func (d *db) SetStatus(ctx context.Context, status string) error {
	return d.setStatusQuery(ctx, status)
}

func (d *db) StatusList(ctx context.Context, filter *models.Filter) ([]*desc.StatusListReply_Action, error) {
	return d.statusListQuery(ctx, filter)
}
