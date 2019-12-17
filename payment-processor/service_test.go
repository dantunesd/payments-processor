package payment

import (
	"context"
	"errors"
	"testing"
)

func TestService_ProcessPayment(t *testing.T) {
	type fields struct {
		r ISourceRepository
	}
	type args struct {
		ctx context.Context
		p   Payment
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
				r: &SourceRepositoryMock{
					getByID: func(ctx context.Context, ID string) (Source, error) {
						return Source{}, nil
					},
				},
			},
			args{
				context.Background(),
				Payment{},
			},
			false,
		},
		{
			"invalid source",
			fields{
				r: &SourceRepositoryMock{
					getByID: func(ctx context.Context, ID string) (Source, error) {
						return Source{}, errors.New("Invalid source_id")
					},
				},
			},
			args{
				context.Background(),
				Payment{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r)

			if err := s.ProcessPayment(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Service.ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type SourceRepositoryMock struct {
	getByID func(ctx context.Context, ID string) (Source, error)
}

func (s *SourceRepositoryMock) GetByID(ctx context.Context, ID string) (Source, error) {
	return s.getByID(ctx, ID)
}
