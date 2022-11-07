package service

type CounterServiceInterface interface {
	Get() (int64, error)
	Increment(userAgent string) (int64, error)
}
