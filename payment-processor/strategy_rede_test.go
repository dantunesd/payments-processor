package payment

import (
	"context"
	"errors"
	"testing"
)

func TestRedeStrategy_Process(t *testing.T) {
	type fields struct {
		r IRedeRepository
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
				r: redeRepositoryMock{
					func(context.Context, RedeRequestBody) (ITransaction, error) {
						return &redeTransactionMock{
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
			"trasaction failed",
			fields{
				r: redeRepositoryMock{
					func(context.Context, RedeRequestBody) (ITransaction, error) {
						return &redeTransactionMock{
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
				r: redeRepositoryMock{
					func(context.Context, RedeRequestBody) (ITransaction, error) {
						return nil, errors.New("failed")
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
			c := NewRedeStrategy(tt.fields.r)
			if err := c.Process(tt.args.ctx, tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RedeStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type redeRepositoryMock struct {
	transaction func(context.Context, RedeRequestBody) (ITransaction, error)
}

func (r redeRepositoryMock) Transaction(ctx context.Context, rrb RedeRequestBody) (ITransaction, error) {
	return r.transaction(ctx, rrb)
}

type redeTransactionMock struct {
	paymentSucceeded func() error
}

func (r redeTransactionMock) PaymentSucceeded() error {
	return r.paymentSucceeded()
}
