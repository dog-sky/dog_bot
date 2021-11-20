package app

import (
	"context"
	"log"
	"net"

	"github.com/dog-sky/dog_bot/internal/config"
	"github.com/dog-sky/dog_bot/internal/dog/interact"

	"google.golang.org/grpc"
)

type App struct {
	ctx context.Context
	cfg *config.Config
	s   *grpc.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	s := grpc.NewServer()

	interact.New(s)

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
