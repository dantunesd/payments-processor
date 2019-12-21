package payment

import (
	"errors"
	"reflect"
	"testing"
)

func TestAcquirerProvider_GetAcquirer(t *testing.T) {
	type fields struct {
		Acquirers map[AcquirerStrategy]IAcquirerStrategy
	}
	type args struct {
		as AcquirerStrategy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IAcquirerStrategy
	}{
		{
			"return an acquirer with success",
			fields{
				Acquirers: map[AcquirerStrategy]IAcquirerStrategy{
					Cielo: CieloStrategy{},
					Rede:  RedeStrategy{},
				},
			},
			args{
				Cielo,
			},
			CieloStrategy{},
		},
		{
			"return an inexistent acquirer",
			fields{
				Acquirers: map[AcquirerStrategy]IAcquirerStrategy{
					Cielo: CieloStrategy{},
					Rede:  RedeStrategy{},
				},
			},
			args{
				"undefined",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := NewAcquirerProvider(tt.fields.Acquirers)
			if got := ap.GetAcquirer(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AcquirerProvider.GetAcquirer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCieloStrategy_Process(t *testing.T) {
	type fields struct {
		r ICieloRepository
	}
	type args struct {
		p Payment
		s Source
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
					func(CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{}, nil
					},
				},
			},
			args{
				Payment{},
				Source{},
			},
			false,
		},
		{
			"failed to process",
			fields{
				r: repositoryMock{
					func(CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{}, errors.New("failed to request")
					},
				},
			},
			args{
				Payment{},
				Source{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloStrategy(tt.fields.r)
			if err := c.Process(tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CieloStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedeStrategy_Process(t *testing.T) {
	type args struct {
		p Payment
		s Source
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"process with success",
			args{
				Payment{},
				Source{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedeStrategy()
			if err := c.Process(tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RedeStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type repositoryMock struct {
	sale func(CieloRequestBody) (*CieloResponseBody, error)
}

func (r repositoryMock) Sale(c CieloRequestBody) (*CieloResponseBody, error) {
	return r.sale(c)
}
