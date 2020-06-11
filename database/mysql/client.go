package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	// Client struct
	Client struct {
		*sqlx.DB
		logMode bool
		context context.Context
	}
)

// WithContext add context to sql
func (db *Client) WithContext(ctx context.Context) *Client {
	return &Client{
		DB:      db.DB,
		logMode: db.logMode,
		context: ctx,
	}
}

// Query queries the database and returns an *sql.Rows.
func (db *Client) Query(query string, args ...interface{}) (*sql.Rows, error) {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.QueryContext(db.context, query, args...)
	}

	return db.DB.Query(query, args...)
}

// QueryRow queries the database and returns an *sqlx.Row.
func (db *Client) QueryRow(query string, args ...interface{}) *sql.Row {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.QueryRowContext(db.context, query, args...)
	}
	return db.DB.QueryRow(query, args...)
}

// Queryx queries the databas	e and returns an *sqlx.Rows.
func (db *Client) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.QueryxContext(db.context, query, args...)
	}

	return db.DB.Queryx(query, args...)
}

// QueryRowx queries the database and returns an *sqlx.Row.
func (db *Client) QueryRowx(query string, args ...interface{}) *sqlx.Row {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.QueryRowxContext(db.context, query, args...)
	}
	return db.DB.QueryRowx(query, args...)
}

// Exec using master db
func (db *Client) Exec(query string, args ...interface{}) (sql.Result, error) {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.ExecContext(db.context, query, args...)
	}
	return db.DB.Exec(query, args...)
}

// Select using slave db.
func (db *Client) Select(dest interface{}, query string, args ...interface{}) error {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.SelectContext(db.context, dest, query, args...)
	}

	return db.DB.Select(dest, query, args...)
}

// Get using slave db.
func (db *Client) Get(dest interface{}, query string, args ...interface{}) error {

	Log(db.logMode, query, args...)
	if db.context != nil {
		return db.DB.GetContext(db.context, dest, query, args...)
	}

	return db.DB.Get(dest, query, args...)
}

// MustBegin starts a transaction, and panics on error.
func (db *Client) MustBegin() *Tx {
	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}

	return &Tx{Tx: tx, logMode: db.logMode, context: db.context}
}

// Begin begins a transaction and returns an *Tx instead of an *sql.Tx.
func (db *Client) Begin() (*Tx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx, logMode: db.logMode, context: db.context}, nil
}
