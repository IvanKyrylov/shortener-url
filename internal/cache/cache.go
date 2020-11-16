package cache

import (
	"sync"
	"time"
)

type ItemCache struct {
	Content    []byte
	Expiration int64
}

type Storage struct {
	items map[string]ItemCache
	mu    *sync.RWMutex
}

func (item ItemCache) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]ItemCache),
		mu:    &sync.RWMutex{},
	}
}

func (s Storage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	if item.Expired() {
		delete(s.items, key)
		return nil
	}
	return item.Content
}

func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = ItemCache{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}
