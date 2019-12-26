package payment

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestRedeRepository_Transaction(t *testing.T) {
	type fields struct {
		r IHTTPRequester
	}
	type args struct {
		ctx context.Context
		rrb RedeRequestBody
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ITransaction
		wantErr bool
	}{
		{
			"request with success",
			fields{
				r: requesterMock{
					post: func(ctx context.Context, path string, body interface{}) (IResponser, error) {
						return &Response{200, []byte{}}, nil
					},
				},
			},
			args{
				context.Background(),
				RedeRequestBody{},
			},
			&RedeTransaction{&Response{200, []byte{}}},
			false,
		},
		{
			"request with error",
			fields{
				r: requesterMock{
					post: func(ctx context.Context, path string, body interface{}) (IResponser, error) {
						return nil, errors.New("failed to request")
					},
				},
			},
			args{
				context.Background(),
				RedeRequestBody{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedeRepository(tt.fields.r)
			got, err := c.Transaction(tt.args.ctx, tt.args.rrb)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedeRepository.Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedeRepository.Transaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
