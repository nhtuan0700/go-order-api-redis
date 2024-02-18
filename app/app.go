package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	config Config
	rdb    *redis.Client
}

func NewApp(config Config) *App {
	app := &App{
		config: config,
		rdb:    NewRedis(config),
	}

	return app
}

func Start(ctx context.Context) error {
	app := NewApp(LoadConfig())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.ServerPort),
	}
	app.loadRoutes(server)

	// ch is a chanel that is used for send and received error in Start and Stop server
	ch := make(chan error, 1)
	go StartServer(server, ch)

	StartRedis(ctx, app.rdb)
	defer StopRedis(app.rdb)

	return StopServer(ctx, server, ch)
}

func StartServer(server *http.Server, ch chan<- error) {
	log.Println("Starting server!")
	err := server.ListenAndServe()
	if err != nil {
		ch <- fmt.Errorf("failed to start server: %w", err)
	}
}

func StopServer(ctx context.Context, server *http.Server, ch <- chan error) error {
	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		log.Println("Cancelling")
		// Waiting 5 seconds before shutdown server
		<-timeout.Done()
		log.Println("Server exited")

		return server.Shutdown(timeout)
	}
}
