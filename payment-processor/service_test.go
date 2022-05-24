package payment

import (
	"context"
	"errors"
	"testing"
)

func TestService_ProcessPayment(t *testing.T) {
	type fields struct {
		r ISourceRepository
		a IAcquirerProvider
	}
	type args struct {
		ctx context.Context
		p   Payment
		a   Acquirer
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
				a: &AcquirerProviderMock{
					getAcquirerStrategy: func(ac Acquirer) IAcquirerStrategy {
						return &AcquirerStrategyMock{
							process: func(ctx context.Context, p Payment, s Source) error {
								return nil
							},
						}
					},
				},
			},
			args{
				context.Background(),
				Payment{
					"o123456789",
					Customer{"test"},
					Details{
						Card{"test", "test", 2020, 12},
						100,
						"credit",
						1,
						[]string{"test"},
					},
					Establishment{"test", "test", 12345678},
				},
				Cielo,
			},
			false,
		},
		{
			"invalid payment",
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
				Cielo,
			},
			true,
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
				Payment{
					"o123456789",
					Customer{"test"},
					Details{
						Card{"test", "test", 2020, 12},
						100,
						"credit",
						1,
						[]string{"test"},
					},
					Establishment{"test", "test", 12345678},
				},
				Cielo,
			},
			true,
		},
		{
			"fail to process",
			fields{
				r: &SourceRepositoryMock{
					getByID: func(ctx context.Context, ID string) (Source, error) {
						return Source{}, nil
					},
				},
				a: &AcquirerProviderMock{
					getAcquirerStrategy: func(ac Acquirer) IAcquirerStrategy {
						return &AcquirerStrategyMock{
							process: func(ctx context.Context, p Payment, s Source) error {
								return errors.New("Failed to process")
							},
						}
					},
				},
			},
			args{
				context.Background(),
				Payment{
					"o123456789",
					Customer{"test"},
					Details{
						Card{"test", "test", 2020, 12},
						100,
						"credit",
						1,
						[]string{"test"},
					},
					Establishment{"test", "test", 12345678},
				},
				Cielo,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r, tt.fields.a)

			if err := s.ProcessPayment(tt.args.ctx, tt.args.p, tt.args.a); (err != nil) != tt.wantErr {
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

type AcquirerProviderMock struct {
	getAcquirerStrategy func(Acquirer) IAcquirerStrategy
}

func (a *AcquirerProviderMock) GetAcquirerStrategy(ac Acquirer) IAcquirerStrategy {
	return a.getAcquirerStrategy(ac)
}

type AcquirerStrategyMock struct {
	process func(context.Context, Payment, Source) error
}

func (a *AcquirerStrategyMock) Process(ctx context.Context, p Payment, s Source) error {
	return a.process(ctx, p, s)
}
