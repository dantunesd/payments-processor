package payment

import (
	"testing"
)

func TestRedeTransaction_PaymentSucceeded(t *testing.T) {
	type fields struct {
		r IResponser
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"transaction succeeded",
			fields{
				&redeResponser{
					getStatusCode: func() int {
						return 200
					},
					getBody: func() []byte {
						return []byte(`{"returnMessage":"Successful.","returnCode":"00"}`)
					},
				},
			},
			false,
		},
		{
			"transaction with integration error",
			fields{
				&redeResponser{
					getStatusCode: func() int {
						return 400
					},
					getBody: func() []byte {
						return []byte(`{"returnCode":"37","returnMessage":"CardNumber: Invalid parameter format."}`)
					},
				},
			},
			true,
		},
		{
			"transaction with business error",
			fields{
				&redeResponser{
					getStatusCode: func() int {
						return 412
					},
					getBody: func() []byte {
						return []byte(`{"returnCode":"42","returnMessage":"Reference: Order number already exists."}`)
					},
				},
			},
			true,
		},
		{
			"transaction with emissor error",
			fields{
				&redeResponser{
					getStatusCode: func() int {
						return 400
					},
					getBody: func() []byte {
						return []byte(`{"returnCode":"101","returnMessage":"Unauthorized. Problems on the card, contact the issuer."}`)
					},
				},
			},
			true,
		},
		{
			"transaction with server error",
			fields{
				&redeResponser{
					getStatusCode: func() int {
						return 500
					},
					getBody: func() []byte {
						return []byte(`internal server error`)
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RedeTransaction{
				r: tt.fields.r,
			}
			if err := r.PaymentSucceeded(); (err != nil) != tt.wantErr {
				t.Errorf("RedeTransaction.PaymentSucceeded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type redeResponser struct {
	getStatusCode func() int
	getBody       func() []byte
}

func (r *redeResponser) GetStatusCode() int {
	return r.getStatusCode()
}

// GetBody return body
func (r *redeResponser) GetBody() []byte {
	return r.getBody()
}
