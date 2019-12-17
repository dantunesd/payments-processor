package payment

import "testing"

func TestPayment_IsValid(t *testing.T) {
	type fields struct {
		Customer      Customer
		Card          Card
		Sale          Sale
		Establishment Establishment
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"With all fields",
			fields{
				Customer{"test"},
				Card{"test", "test", "test"},
				Sale{100, 1, []string{"test"}},
				Establishment{"test", "test", "test"},
			},
			false,
		},
		{
			"With invalid amount",
			fields{
				Customer{"test"},
				Card{"test", "test", "test"},
				Sale{0, 1, []string{"test"}},
				Establishment{"test", "test", "test"},
			},
			true,
		},
		{
			"With invalid installments",
			fields{
				Customer{"test"},
				Card{"test", "test", "test"},
				Sale{100, 0, []string{"test"}},
				Establishment{"test", "test", "test"},
			},
			true,
		},
		{
			"With invalid itens",
			fields{
				Customer{"test"},
				Card{"test", "test", "test"},
				Sale{100, 1, []string{}},
				Establishment{"test", "test", "test"},
			},
			true,
		},
		{
			"missing some fields",
			fields{
				Customer: Customer{
					Name: "test",
				},
			},
			true,
		},
		{
			"missing all fields",
			fields{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Payment{
				Customer:      tt.fields.Customer,
				Card:          tt.fields.Card,
				Sale:          tt.fields.Sale,
				Establishment: tt.fields.Establishment,
			}
			if err := p.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("Payment.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
