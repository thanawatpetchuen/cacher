package cacher

import (
	"errors"
	"log"
	"time"

	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	gocache "github.com/patrickmn/go-cache"
)

var ErrNotFound = errors.New("Value not found in GoCache store")

type Cacher interface {
	Do(key string, f interface{}, params ...interface{}) (interface{}, error)
	Engine() *cache.Cache
}

type cacher struct {
	cacheEngine *cache.Cache
}

func (c cacher) Do(key string, f interface{}, params ...interface{}) (interface{}, error) {
	result, err := c.cacheEngine.Get(key)
	if err != nil && err.Error() != ErrNotFound.Error() {
		return nil, err
	}
	if result == nil {
		log.Println("Set cache with key:", key)
		return f, c.cacheEngine.Set(key, f, &store.Options{})
	} else {
		return result, nil
	}
}

func (c cacher) Engine() *cache.Cache {
	return c.cacheEngine
}

func New() Cacher {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := store.NewGoCache(gocacheClient, nil)

	cacheManager := cache.New(gocacheStore)

	return cacher{cacheEngine: cacheManager}
}
