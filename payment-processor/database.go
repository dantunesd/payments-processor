package payment

import (
	"context"
)

// IScanner interface for scanner
type IScanner interface {
	Scan(dest ...interface{}) error
}

// IQuerier interface for db queries
type IQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner
}
