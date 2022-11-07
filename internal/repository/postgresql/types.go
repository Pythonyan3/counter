package postgresql

type counter struct {
	Id      int64 `db:"id"`
	Counter int64 `db:"counter"`
}
