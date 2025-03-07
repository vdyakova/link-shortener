package shortener

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

//go:generate mockgen -source=shortener.go -destination=mocks/mock_storage.go -package=mocks

type cache interface {
	Save(ctx context.Context, dataSh string, dataLg string) error
	HasData(ctx context.Context, data string) (bool, error)
	FindLongURL(ctx context.Context, dataSh string) (string, error)
	GiveInfoCache(ctx context.Context) ([]string, error)
}

type ShortenerService interface {
	ShortLink(ctx context.Context, data string) (string, error)
	LongLink(ctx context.Context, data string) (string, error)
	GiveInfoAboutLink(ctx context.Context) ([]string, error)
}

type shortenerSrv struct {
	linkCache cache
}

func New(cache cache) *shortenerSrv {
	return &shortenerSrv{linkCache: cache}
}

func (l *shortenerSrv) ShortLink(ctx context.Context, data string) (string, error) {
	exists, err := l.linkCache.HasData(ctx, data)
	if err != nil {
		return "", err
	}

	if exists {
		return "URL-long and URL-sh already exists", nil
	}

	newLink, err := shortener(data)
	if err != nil {
		return "", err
	}

	err = l.linkCache.Save(ctx, newLink, data)
	if err != nil {
		return "", fmt.Errorf("failed in cashe: %w", err)
	}

	return newLink, nil
}

func (l *shortenerSrv) LongLink(ctx context.Context, data string) (string, error) {
	longLink, err := l.linkCache.FindLongURL(ctx, data)
	if err != nil {
		return "", fmt.Errorf("failed in cashe: %w", err)
	}

	return longLink, nil
}

func (l *shortenerSrv) GiveInfoAboutLink(ctx context.Context) ([]string, error) {
	links, err := l.linkCache.GiveInfoCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed in cashe: %w", err)
	}

	return links, nil
}

// shortener - функция сокращения ссылки.
func shortener(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty input url")
	}

	hash := sha256.Sum256([]byte(data))
	encoded := base64.RawURLEncoding.EncodeToString(hash[:6])
	shortUrl := strings.ReplaceAll(encoded, "/", "_")

	return "https://short.ly/" + shortUrl, nil
}
