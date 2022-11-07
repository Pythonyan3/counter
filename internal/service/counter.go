package service

import "github.com/Pythonyan3/Counter/internal/repository"

// verify interface compliance.
var _ CounterServiceInterface = (*CounterService)(nil)

type CounterService struct {
	repo repository.CounterRepositoryInterface
}

// NewCounterService CounterService constructor function.
func NewCounterService(repo repository.CounterRepositoryInterface) *CounterService {
	return &CounterService{repo: repo}
}

// Get return counter current value.
func (service *CounterService) Get() (int64, error) {
	return service.repo.Get()
}

// Increment perform counter value increment by 1 and return new value.
func (service *CounterService) Increment(additional string) (int64, error) {
	return service.repo.Increment(additional)
}
