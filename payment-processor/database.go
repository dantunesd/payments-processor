package payment

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
)

// IScanner interface for db scanner.
type IScanner interface {
	Scan(dest ...interface{}) error
}

// IQuerier interface for db queries.
type IQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner
}

// LoggableDBWrapper is loggable db wrapper.
type LoggableDBWrapper struct {
	db *sql.DB
	lg *zap.Logger
}

// NewLoggableDBWrapper LoggableDBWrapper's constructor.
func NewLoggableDBWrapper(db *sql.DB, lg *zap.Logger) *LoggableDBWrapper {
	return &LoggableDBWrapper{
		db: db,
		lg: lg,
	}
}

// QueryRowContext is a wrapper for sql DB QueryRowContext.
func (l *LoggableDBWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner {

	l.lg.Info(
		"logging query",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return l.db.QueryRowContext(ctx, query, args...)
}
