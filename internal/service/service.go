package service

type CounterServiceInterface interface {
	Get() int64
	Increment() (int64, error)
}
