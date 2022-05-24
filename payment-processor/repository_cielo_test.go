package payment

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestCieloRepository_Transaction(t *testing.T) {
	type fields struct {
		r IHTTPRequester
	}
	type args struct {
		ctx context.Context
		crb CieloRequestBody
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
				CieloRequestBody{},
			},
			&CieloTransaction{&Response{200, []byte{}}},
			false,
		},
		{
			"fail to request",
			fields{
				r: requesterMock{
					post: func(ctx context.Context, path string, body interface{}) (IResponser, error) {
						return nil, errors.New("failed for some reason")
					},
				},
			},
			args{
				context.Background(),
				CieloRequestBody{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloRepository(tt.fields.r)
			got, err := c.Transaction(tt.args.ctx, tt.args.crb)
			if (err != nil) != tt.wantErr {
				t.Errorf("CieloRepository.Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CieloRepository.Transaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

type requesterMock struct {
	post func(ctx context.Context, path string, body interface{}) (IResponser, error)
}

func (r requesterMock) Post(ctx context.Context, path string, body interface{}) (IResponser, error) {
	return r.post(ctx, path, body)
}
