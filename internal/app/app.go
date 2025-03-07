package app

import (
	"context"
	"github.com/gorilla/mux"
	"linkSh/internal/cache"
	"linkSh/internal/httpClient"
	"linkSh/internal/httpClient/middleware"
	"linkSh/internal/shortener"
	"log"
	"net/http"
)

type App struct {
	httpClient *httpClient.Client
}

func NewApp() *App {
	serviceSh := shortener.New(cache.NewCache())

	return &App{
		httpClient: httpClient.New(serviceSh),
	}
}

func (a *App) Run(ctx context.Context) error {
	r := mux.NewRouter()
	//создание новой короткой ссылки из длинной
	r.Handle("/api/shorten", middleware.WithContext(ctx, http.HandlerFunc(a.httpClient.ShortenerLink))).Methods("POST")
	// выдача длиной ссылки из короткой
	r.Handle("/api/shorten", middleware.WithContext(ctx, http.HandlerFunc(a.httpClient.OriginalLink))).Methods("GET")

	log.Println("Server started at http://localhost:8080")

	return http.ListenAndServe(":8080", r)
}
