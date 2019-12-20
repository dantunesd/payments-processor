package payment

import (
	"context"
	"database/sql"
)

// IScanner interface for db scanner
type IScanner interface {
	Scan(dest ...interface{}) error
}

// IQuerier interface for db queries
type IQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner
}

// DBWrapper wrapper for DB
type DBWrapper struct {
	DB *sql.DB
}

// QueryRowContext wrapper for QueryRowContext
func (q *DBWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner {
	return q.DB.QueryRowContext(ctx, query, args...)
}
