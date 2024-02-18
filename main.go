package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/nhtuan0700/orders-api/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}
