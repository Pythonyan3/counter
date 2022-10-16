package repository

import "github.com/Pythonyan3/Counter/internal/repository/inmemory"

// verify interface compliance for specific structs.
var _ CounterRepositoryInterface = (*inmemory.InMemoryRepository)(nil)

type CounterRepositoryInterface interface {
	Get() int64
	Increment() (int64, error)
}
