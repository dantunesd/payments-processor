package payment

import (
	"context"
	"database/sql"
)

// IQuerier interface for db queries
type IQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
