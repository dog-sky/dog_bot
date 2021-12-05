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

	shutDownFns []func()
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	s := grpc.NewServer()

	dbService, err := db.New(ctx, cfg.DbURL)
	if err != nil {
		return nil, err
	}

	interactSerivice := interact.New(dbService)
	interactDesc.RegisterDogServer(s, interactSerivice)

	shutDownFns := []func(){
		s.GracefulStop,
		dbService.ShutDown,
	}

	return &App{
		ctx:         ctx,
		cfg:         cfg,
		s:           s,
		shutDownFns: shutDownFns,
	}, nil
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

func (a *App) ShutDown() {
	log.Println("shutting down")

	for _, fn := range a.shutDownFns {
		fn()
	}
}
