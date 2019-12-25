package payment

// Transaction is a interface for a transaction
type Transaction interface {
	IsSucceeded() bool
	GetError() error
}
