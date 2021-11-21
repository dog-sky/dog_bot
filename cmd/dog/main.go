package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dog-sky/dog_bot/internal/app"
	"github.com/dog-sky/dog_bot/internal/config"
	"golang.org/x/sync/errgroup"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading conf: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	g, ctx := errgroup.WithContext(ctx)
	service, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("err creating servise: %s", err.Error())
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	go func() {
		<-ctx.Done()
		service.ShutDown()
	}()

	g.Go(func() error {
		return service.Serve()
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}
}
