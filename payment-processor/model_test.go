package payment

import "testing"

func TestPayment_IsValid(t *testing.T) {
	type fields struct {
		OrderID       string
		Customer      Customer
		Details       Details
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
			false,
		},
		{
			"With invalid amount",
			fields{
				"o123456789",
				Customer{"test"},
				Details{
					Card{"test", "test", 2020, 12},
					99,
					"credit",
					1,
					[]string{"test"},
				},
				Establishment{"test", "test", 12345678},
			},
			true,
		},
		{
			"With invalid installments",
			fields{
				"o123456789",
				Customer{"test"},
				Details{
					Card{"test", "test", 2020, 12},
					100,
					"credit",
					0,
					[]string{"test"},
				},
				Establishment{"test", "test", 12345678},
			},
			true,
		},
		{
			"With invalid itens",
			fields{
				"o123456789",
				Customer{"test"},
				Details{
					Card{"test", "test", 2020, 12},
					100,
					"credit",
					1,
					[]string{},
				},
				Establishment{"test", "test", 12345678},
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
				OrderID:       tt.fields.OrderID,
				Customer:      tt.fields.Customer,
				Details:       tt.fields.Details,
				Establishment: tt.fields.Establishment,
			}
			if err := p.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("Payment.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
