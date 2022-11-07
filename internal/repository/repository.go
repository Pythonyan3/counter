package repository

import (
	"github.com/Pythonyan3/Counter/internal/repository/inmemory"
	"github.com/Pythonyan3/Counter/internal/repository/postgresql"
)

// verify interface compliance for specific structs.
var _ CounterRepositoryInterface = (*inmemory.InMemoryRepository)(nil)
var _ CounterRepositoryInterface = (*postgresql.PostgresRepository)(nil)

type CounterRepositoryInterface interface {
	Get() (int64, error)
	Increment(additional string) (int64, error)
}
