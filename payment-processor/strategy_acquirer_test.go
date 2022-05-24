package payment

import (
	"reflect"
	"testing"
)

func TestAcquirerProvider_GetAcquirer(t *testing.T) {
	type fields struct {
		Acquirers AcquirerStrategies
	}
	type args struct {
		as Acquirer
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
				Acquirers: AcquirerStrategies{
					Cielo: &CieloStrategy{},
					Rede:  &RedeStrategy{},
				},
			},
			args{
				Cielo,
			},
			&CieloStrategy{},
		},
		{
			"return an inexistent acquirer",
			fields{
				Acquirers: AcquirerStrategies{
					Cielo: &CieloStrategy{},
					Rede:  &RedeStrategy{},
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
			if got := ap.GetAcquirerStrategy(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AcquirerProvider.GetAcquirer() = %v, want %v", got, tt.want)
			}
		})
	}
}
