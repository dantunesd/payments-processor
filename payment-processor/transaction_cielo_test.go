package payment

import (
	"testing"
)

func TestCieloTransaction_IsSucceeded(t *testing.T) {
	type fields struct {
		r IResponser
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"transaction succeeded",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 200
					},
					getBody: func() []byte {
						return []byte(`{"Payment":{"Status":2,"ReturnMessage":"Operation Successful","ReturnCode":"4"}}`)
					},
				},
			},
			true,
		},
		{
			"transaction with emissor error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 200
					},
					getBody: func() []byte {
						return []byte(`{"Payment":{"Status":3,"ReturnMessage":"Not Authorized","ReturnCode":"5"}}`)
					},
				},
			},
			false,
		},
		{
			"transaction with integration error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 400
					},
					getBody: func() []byte {
						return []byte(`{}`)
					},
				},
			},
			false,
		},
		{
			"transaction with server error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 500
					},
					getBody: func() []byte {
						return []byte(`internal server error`)
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloTransaction(tt.fields.r)
			if got := c.IsSucceeded(); got != tt.want {
				t.Errorf("CieloTransaction.IsSucceeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCieloTransaction_GetError(t *testing.T) {
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
				&cieloResponser{
					getStatusCode: func() int {
						return 200
					},
					getBody: func() []byte {
						return []byte(`{"Payment":{"Status":2,"ReturnMessage":"Operation Successful","ReturnCode":"4"}}`)
					},
				},
			},
			false,
		},
		{
			"transaction with emissor error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 200
					},
					getBody: func() []byte {
						return []byte(`{"Payment":{"Status":3,"ReturnMessage":"Not Authorized","ReturnCode":"5"}}`)
					},
				},
			},
			true,
		},
		{
			"transaction with integration error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 400
					},
					getBody: func() []byte {
						return []byte(`{}`)
					},
				},
			},
			true,
		},
		{
			"transaction with server error",
			fields{
				&cieloResponser{
					getStatusCode: func() int {
						return 500
					},
					getBody: func() []byte {
						return []byte(`{}`)
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CieloTransaction{
				r: tt.fields.r,
			}
			if err := c.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("CieloTransaction.GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type cieloResponser struct {
	getStatusCode func() int
	getBody       func() []byte
}

func (r *cieloResponser) GetStatusCode() int {
	return r.getStatusCode()
}

// GetBody return body
func (r *cieloResponser) GetBody() []byte {
	return r.getBody()
}
