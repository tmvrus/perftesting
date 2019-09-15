package kvstorage

import "sync"

type (
	kvStorage struct {
		sync.RWMutex
		data map[int]int
	}
)

func NewSingleStorage() *kvStorage {
	return &kvStorage{
		data: map[int]int{},
	}
}

func (s *kvStorage) Get(key int) int {
	s.RLock()
	defer s.RUnlock()
	return s.data[key]
}

func (s *kvStorage) Set(key, value int) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}
