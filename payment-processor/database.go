package payment

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
)

// IScanner interface for db scanner
type IScanner interface {
	Scan(dest ...interface{}) error
}

// IQuerier interface for db queries
type IQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner
}

// LoggableDBWrapper wrapper for DB
type LoggableDBWrapper struct {
	DB     *sql.DB
	Logger *zap.Logger
}

// QueryRowContext wrapper for QueryRowContext
func (q *LoggableDBWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner {

	q.Logger.Info(
		"logging queryrow",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return q.DB.QueryRowContext(ctx, query, args...)
}
