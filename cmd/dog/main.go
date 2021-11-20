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

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)
	service := app.New(ctx, cfg)

	g.Go(func() error {
		<-gCtx.Done()
		return service.ShutDown()
	})

	g.Go(func() error {
		return service.Serve()
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}

	log.Println("shutdown signal")
}
