package kvstorage

import "sync"

const (
	ShardCount = 12
)

type (
	ShardedStorage []*kvStorage
	ShardedSyncMap []*sync.Map
)

func NewShardedStorage() ShardedStorage {
	s := make(ShardedStorage, ShardCount)
	for i := 0; i < ShardCount; i++ {
		s[i] = NewSingleStorage()
	}

	return s
}

func (s ShardedStorage) shardNum(key int) int {
	return key % ShardCount
}

func (s ShardedStorage) Get(key int) int {
	return s[s.shardNum(key)].Get(key)
}

func (s ShardedStorage) Set(key, value int) {
	s[s.shardNum(key)].Set(key, value)
}

func NewShardedSyncMap() ShardedSyncMap {
	s := make(ShardedSyncMap, ShardCount)
	for i := 0; i < ShardCount; i++ {
		s[i] = &sync.Map{}
	}

	return s
}

func (s ShardedSyncMap) shardNum(key int) int {
	return key % ShardCount
}

func (s ShardedSyncMap) Get(key int) (interface{}, bool) {
	return s[s.shardNum(key)].Load(key)
}

func (s ShardedSyncMap) Set(key, value int) {
	s[s.shardNum(key)].Store(key, value)
}
