package payment

import (
	"context"
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
					transaction: func(context.Context, RedeRequestBody) (*RedeResponseBody, error) {
						return &RedeResponseBody{ReturnCode: "00"}, nil
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
	transaction func(context.Context, RedeRequestBody) (*RedeResponseBody, error)
}

func (r redeRepositoryMock) Transaction(ctx context.Context, rrb RedeRequestBody) (*RedeResponseBody, error) {
	return r.transaction(ctx, rrb)
}
