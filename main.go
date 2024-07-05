package main

import (
	"context"
	"log"

	"booking-service/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("business_errors while creating app: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("business_errors while running app: %s", err.Error())
	}
	defer app.Shutdown()
}
