package inmemory

import "sync/atomic"

type InMemoryRepository struct {
	counter int64
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

// Get return current counter value.
func (repository *InMemoryRepository) Get() (int64, error) {
	return repository.counter, nil
}

// Increment perform counter value increment by 1 and return new value.
func (repository *InMemoryRepository) Increment(_ string) (int64, error) {
	return atomic.AddInt64(&repository.counter, 1), nil
}
