package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"linkSh/internal/cashe"
	"strings"
)
//go:generate mockgen -source=shortener.go -destination=mocks/mock_storage.go -package=mocks

type ShortenerService interface {
	ShortLink(ctx context.Context, data string) (string, error)
	LongLink(ctx context.Context, data string) (string, error)
	GiveInfoAboutLink(ctx context.Context) ([]string, error)
}

type shortenersrv struct {
	linkcache cashe.Cache
}

func NewShortener(cache cashe.Cache) ShortenerService {
	return &shortenersrv{linkcache: cache}
}
func (l *shortenersrv) ShortLink(ctx context.Context, data string) (string, error) {
	exists, err := l.linkcache.HasData(ctx, data)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("url exists")
	}
	newLink, err := Shortener(data)
	if err != nil {
		return "", err
	}
	err = l.linkcache.Save(ctx, newLink, data)
	if err != nil {
		return "", err
	}
	return newLink, nil
}
func (l *shortenersrv) LongLink(ctx context.Context, data string) (string, error) {
	longLink, err := l.linkcache.FindLongURL(ctx, data)
	if err != nil {
		return "", errors.New("not find long link")
	}
	return longLink, nil
}
func (l *shortenersrv) GiveInfoAboutLink(ctx context.Context) ([]string, error) {

	links, err := l.linkcache.GiveInfoCache(ctx)
	if err != nil {
		return nil, err
	}
	return links, nil
}

// функция сокращения ссылки

func Shortener(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty input url")
	}
	hash := sha256.Sum256([]byte(data))
	encoded := base64.RawURLEncoding.EncodeToString(hash[:6])
	shortUrl := strings.ReplaceAll(encoded, "/", "_")
	return "https://short.ly/" + shortUrl, nil

}
