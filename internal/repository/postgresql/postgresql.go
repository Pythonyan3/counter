package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	counterTableName string = "counter"
	historyTableName string = "history"
)

type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository PostgresRepository struct constructor func.
func NewPostgresRepository() (*PostgresRepository, error) {
	// so bad to hardcode credentials...
	db, err := sqlx.Open("postgres", "host=counterdb port=5432 user=postgres dbname=counter password=postgres sslmode=disable")

	// check successfully connection
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	transaction, err := db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("db.Beginx: %w", err)
	}

	// perform history table initialize
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id serial PRIMARY KEY, timestamp timestamp NOT NULL, user_agent varchar);", historyTableName)
	_, err = transaction.Queryx(query)
	if err != nil {
		transaction.Rollback()
		return nil, fmt.Errorf("create table %v db.Queryx: %w", historyTableName, err)
	}

	// perform counter table initialize
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id serial PRIMARY KEY, counter integer NOT NULL DEFAULT 0);", counterTableName)
	_, err = transaction.Queryx(query)
	if err != nil {
		transaction.Rollback()
		return nil, fmt.Errorf("create table %v db.Queryx: %w", counterTableName, err)
	}

	var recordsCount int64
	err = transaction.Get(&recordsCount, fmt.Sprintf("SELECT count(*) FROM %v;", counterTableName))
	if err != nil {
		transaction.Rollback()
		return nil, fmt.Errorf("db.Get: %w", err)
	}

	if recordsCount == 0 {
		// perform table initialization
		query := fmt.Sprintf("INSERT INTO %v (counter) VALUES (0);", counterTableName)
		_, err = transaction.Queryx(query)
		if err != nil {
			transaction.Rollback()
			return nil, fmt.Errorf("init row insert db.Queryx: %w", err)
		}
	}

	transaction.Commit()

	return &PostgresRepository{db: db}, nil
}

// Close perform db connection closing.
func (repository *PostgresRepository) Close() error {
	return repository.db.Close()
}

// Get return current counter value.
func (repository *PostgresRepository) Get() (int64, error) {
	var cntr counter = counter{}

	// build query string
	query := fmt.Sprintf("SELECT counter FROM %s;", counterTableName)

	// evalate select query and parse new row data to counter struct
	row := repository.db.QueryRowx(query)
	err := row.StructScan(&cntr)
	if err != nil {
		return 0, fmt.Errorf("row.StructScan: %w", err)
	}

	return cntr.Counter, nil
}

// Increment perform counter value increment by 1 and return new value.
func (repository *PostgresRepository) Increment(userAgent string) (int64, error) {
	var cntr counter = counter{}

	// start new db transaction
	transaction, err := repository.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("repository.db.Beginx: %w", err)
	}

	// build query string
	query := fmt.Sprintf("UPDATE %v set counter=counter+1 RETURNING counter;", counterTableName)

	// perform counter increment.
	row := transaction.QueryRowx(query)
	err = row.StructScan(&cntr)
	if err != nil {
		transaction.Rollback()
		return 0, fmt.Errorf("row.StructScan: %w", err)
	}

	// build query string
	query = fmt.Sprintf("INSERT INTO %v (timestamp, user_agent) VALUES (now(), $1);", historyTableName)

	// write record to history table.
	_, err = transaction.Queryx(query, userAgent)
	if err != nil {
		transaction.Rollback()
		return 0, fmt.Errorf("transaction.Queryx: %w", err)
	}

	err = transaction.Commit()
	if err != nil {
		return 0, fmt.Errorf("transaction.Commit: %w", err)
	}

	return cntr.Counter, nil
}
