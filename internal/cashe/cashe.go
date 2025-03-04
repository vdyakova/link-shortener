package cashe

import (
	"context"
	"errors"
	"fmt"
)

// Ключ: короткий Url, Значение: длинный url

type Cache interface {
	Save(ctx context.Context, dataSh string, dataLg string) error
	HasData(ctx context.Context, data string) (bool, error)
	FindLongURL(ctx context.Context, dataSh string) (string, error)
	GiveInfoCache(ctx context.Context) ([]string, error)
}

type storage struct {
	cacheMap map[string]string
}

func NewCache() *storage {
	return &storage{
		cacheMap: make(map[string]string),
	}
}

func (s *storage) Save(ctx context.Context, dataSh string, dataLg string) error {

	if ctx.Err() != nil {
		return errors.New("context canceled")
	}
	if _, exists := s.cacheMap[dataSh]; exists {
		return errors.New("short link already exists")
	}
	s.cacheMap[dataSh] = dataLg
	fmt.Println("Сохранено:", dataSh, "->", dataLg)
	return nil
}
func (s *storage) FindLongURL(ctx context.Context, dataSh string) (string, error) {
	if ctx.Err() != nil {
		return "", errors.New("context canceled")
	}
	val, exists := s.cacheMap[dataSh]
	fmt.Println("find long url", dataSh, val, exists)
	if !exists {
		return "", errors.New("short link not found")
	}
	return val, nil
}

func (s *storage) HasData(ctx context.Context, data string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("context canceled")
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
		return nil, errors.New("context canceled")
	}
	var result []string

	for i, val := range s.cacheMap {
		result = append(result, i, "->", val)
	}
	return result, nil
}
