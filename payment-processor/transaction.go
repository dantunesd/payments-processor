package payment

// ITransaction is a interface for a transaction.
type ITransaction interface {
	PaymentSucceeded() error
}
