package app

import (
	"context"
	"log"
	"net"

	"github.com/dog-sky/dog_bot/internal/config"
	"github.com/dog-sky/dog_bot/internal/dog/interact"
	"github.com/dog-sky/dog_bot/internal/service/db"

	interactDesc "github.com/dog-sky/dog_bot/pkg/dog/api"

	"google.golang.org/grpc"
)

type App struct {
	ctx context.Context
	cfg *config.Config

	db db.DB
	s  *grpc.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	s := grpc.NewServer()
	dbService := db.New()
	interactSerivice := interact.New(dbService)

	interactDesc.RegisterDogServer(s, interactSerivice)

	return &App{
		ctx: ctx,
		cfg: cfg,
		s:   s,
	}
}

func (a *App) Serve() error {
	lis, err := net.Listen("tcp", a.cfg.AppPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

		return err
	}

	log.Printf("server listening at %v", lis.Addr())

	return a.s.Serve(lis)
}

func (a *App) ShutDown() error {
	a.s.GracefulStop()

	return nil
}
