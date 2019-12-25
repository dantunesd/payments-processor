package payment

import (
	"context"
	"errors"
	"testing"
)

func TestCieloStrategy_Process(t *testing.T) {
	type fields struct {
		r ICieloRepository
	}
	type args struct {
		ctx context.Context
		p   Payment
		s   Source
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"process with success",
			fields{
				r: cieloRepositoryMock{
					func(context.Context, CieloRequestBody) (ITransaction, error) {
						return &cieloTransactionMock{
							func() error { return nil },
						}, nil
					},
				},
			},
			args{
				context.Background(),
				Payment{},
				Source{},
			},
			false,
		},
		{
			"transaction failed",
			fields{
				r: cieloRepositoryMock{
					func(context.Context, CieloRequestBody) (ITransaction, error) {
						return &cieloTransactionMock{
							func() error { return errors.New("failed") },
						}, nil
					},
				},
			},
			args{
				context.Background(),
				Payment{},
				Source{},
			},
			true,
		},
		{
			"error",
			fields{
				r: cieloRepositoryMock{
					func(context.Context, CieloRequestBody) (ITransaction, error) {
						return nil, errors.New("error x")
					},
				},
			},
			args{
				context.Background(),
				Payment{},
				Source{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloStrategy(tt.fields.r)
			if err := c.Process(tt.args.ctx, tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CieloStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type cieloRepositoryMock struct {
	transaction func(context.Context, CieloRequestBody) (ITransaction, error)
}

func (c cieloRepositoryMock) Transaction(ctx context.Context, crb CieloRequestBody) (ITransaction, error) {
	return c.transaction(ctx, crb)
}

type cieloTransactionMock struct {
	paymentSucceeded func() error
}

func (c cieloTransactionMock) PaymentSucceeded() error {
	return c.paymentSucceeded()
}
