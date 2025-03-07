package cache

import (
	"context"
	"errors"
	"log"
)

// Ключ: короткий Url, Значение: длинный url

var ErrCtxDone = errors.New("context canceled")

type storage struct {
	cacheMap map[string]string
}

func NewCache() *storage {
	return &storage{
		cacheMap: make(map[string]string), // Начальное значение
	}
}

func (s *storage) Save(ctx context.Context, dataSh string, dataLg string) error {
	if ctx.Err() != nil {
		return ErrCtxDone
	}

	if _, exists := s.cacheMap[dataSh]; exists {
		return errors.New("short link already exists")
	}

	s.cacheMap[dataSh] = dataLg
	log.Println("Сохранено:", dataSh, "->", dataLg)

	return nil
}

func (s *storage) FindLongURL(ctx context.Context, dataSh string) (string, error) {
	if ctx.Err() != nil {
		return "", ErrCtxDone
	}

	val, exists := s.cacheMap[dataSh]
	log.Println("find long url", dataSh, val, exists)
	if !exists {
		return "", errors.New("short link not found")
	}

	return val, nil
}

func (s *storage) HasData(ctx context.Context, data string) (bool, error) {
	if ctx.Err() != nil {
		return false, ErrCtxDone
	}

	for _, link := range s.cacheMap {
		if link == data {
			return true, nil
		}
	}

	return false, nil
}

func (s *storage) GiveInfoCache(ctx context.Context) ([]string, error) {
	if ctx.Err() != nil {
		return nil, ErrCtxDone
	}

	var result []string

	for i, val := range s.cacheMap {
		result = append(result, i, "->", val)
	}

	return result, nil
}
