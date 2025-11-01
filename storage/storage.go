package storage

import "sync"

type Store interface {
	Save(id, url string)
	Get(id string) (string, bool)
}

type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]string)}
}

func (s *MemoryStore) Save(id, url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[id] = url
}

func (s *MemoryStore) Get(id string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RLock()
	val, ok := s.data[id]
	return val, ok
}
