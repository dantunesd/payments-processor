package payment

import (
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
	type args struct {
		p Payment
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
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloStrategy()
			if err := c.Process(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("CieloStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedeStrategy_Process(t *testing.T) {
	type args struct {
		p Payment
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
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedeStrategy()
			if err := c.Process(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("RedeStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
