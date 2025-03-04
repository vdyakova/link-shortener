package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"linkSh/internal/cashe"
	"linkSh/internal/handler"
	"linkSh/internal/middleware"
	"linkSh/internal/service"
	"net/http"
)

type App struct {
	ctx          context.Context
	cacheService cashe.Cache
	serviceSh    service.ShortenerService
	handlerSh    *handler.ShortenerHandler
}

func NewApp(ctx context.Context) (*App, error) {
	cacheService := cashe.NewCache()
	serviceSh := service.NewShortener(cacheService)
	handlerSh := handler.NewShortenerHandler(serviceSh)
	return &App{
		ctx:          ctx,
		cacheService: cacheService,
		serviceSh:    serviceSh,
		handlerSh:    handlerSh,
	}, nil
}

func (a *App) Run() error {
	r := mux.NewRouter()
	//создание новой короткой ссылки из длинной
	r.Handle("/api/shorten", middleware.Withcontext(a.ctx, http.HandlerFunc(a.handlerSh.ShortenerLink))).Methods("POST")
	// выдача длиной ссылки из короткой
	r.Handle("/api/shorten", middleware.Withcontext(a.ctx, http.HandlerFunc(a.handlerSh.OriginalLink))).Methods("GET")
	//вывести содержание кэша
	r.Handle("/api/infoAboutLink", middleware.Withcontext(a.ctx, http.HandlerFunc(a.handlerSh.GiveAboutLink))).Methods("GET")

	fmt.Println("Server started at http://localhost:8080")
	return http.ListenAndServe(":8080", r)
}
