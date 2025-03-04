package main

import (
	"context"
	"linkSh/internal/app"
	"log"
)

func main() {

	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		return
	}
	if err := a.Run(); err != nil {
		log.Fatal("Failed application %s", err.Error())
	}
}
