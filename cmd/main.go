package main

import (
	"context"
	"github.com/vdyakova/link-shortener/internal/app"
	"log"
)

func main() {

	if err := app.NewApp().Run(context.Background()); err != nil {
		log.Fatalf("Failed application %s", err.Error())
	}
}
