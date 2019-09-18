package kvstorage

import (
	"sync"
	"testing"
)

func TestShardedStorage_Set(t *testing.T) {
	s := NewShardedStorage()
	for i := 0; i < ShardCount; i++ {
		s.Set(i, i)
	}

	for i := range s {
		if l := len(s[i].data); l != 1 {
			t.Errorf("shard %d has invalid size %d", i, l)
		}
	}
}

func BenchmarkShardedStorage(b *testing.B) {
	s := NewShardedStorage()
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j <= 100; j++ {
				s.Set(j, i+j)
				_ = s.Get(j)
			}

			wg.Done()
		}(i)

		go func() {
			for j := 0; j <= 100; j++ {
				s.Set(j, j)
				_ = s.Get(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkSingleStorage(b *testing.B) {
	s := NewSingleStorage()
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j <= 100; j++ {
				s.Set(j, i+j)
				_ = s.Get(j)
			}

			wg.Done()
		}(i)

		go func() {
			for j := 0; j <= 100; j++ {
				s.Set(j, j)
				_ = s.Get(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()

}

func BenchmarkSyncMap(b *testing.B) {
	s := sync.Map{}
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j <= 100; j++ {
				s.Store(j, i+j)
				_, _ = s.Load(j)
			}

			wg.Done()
		}(i)

		go func() {
			for j := 0; j <= 100; j++ {
				s.Store(j, j)
				_, _ = s.Load(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()

}

func BenchmarkRefStorage(b *testing.B) {
	data := map[int]*int{}
	for i := 0; i < 100; i++ {
		data[i] = new(int)
	}

	var wg sync.WaitGroup

	wg.Add(b.N)
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j < 100; j++ {
				_ = data[j]
				*data[j] = j + i
			}

			wg.Done()
		}(i)

		go func() {
			for j := 0; j < 100; j++ {
				_ = data[j]
				*data[j] = j + j

			}

			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkShardedSyncMap(b *testing.B) {
	s := NewShardedSyncMap()
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j <= 100; j++ {
				s.Set(j, i+j)
				_, _ = s.Get(j)
			}

			wg.Done()
		}(i)

		go func() {
			for j := 0; j <= 100; j++ {
				s.Set(j, j)
				_, _ = s.Get(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
