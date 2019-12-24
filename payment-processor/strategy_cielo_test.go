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
				r: repositoryMock{
					func(context.Context, CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{
							CieloPaymentResponse{
								Status: 1,
							},
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
			"process with success 2",
			fields{
				r: repositoryMock{
					func(context.Context, CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{
							CieloPaymentResponse{
								Status: 2,
							},
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
			"invalid data",
			fields{
				r: repositoryMock{
					func(context.Context, CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{}, errors.New("failed to request")
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
			"not authorized",
			fields{
				r: repositoryMock{
					func(context.Context, CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{
							CieloPaymentResponse{
								Status: 3,
							},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CieloStrategy{r: tt.fields.r}
			if err := c.Process(tt.args.ctx, tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CieloStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type repositoryMock struct {
	sale func(context.Context, CieloRequestBody) (*CieloResponseBody, error)
}

func (r repositoryMock) Sale(ctx context.Context, c CieloRequestBody) (*CieloResponseBody, error) {
	return r.sale(ctx, c)
}
