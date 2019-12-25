package payment

import (
	"reflect"
	"testing"
)

func TestResponse_GetStatusCode(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"return correctly status code",
			fields{
				200,
				[]byte{},
			},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResponse(tt.fields.StatusCode, tt.fields.Body)
			if got := r.GetStatusCode(); got != tt.want {
				t.Errorf("Response.GetStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetBody(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			"return correctly status code",
			fields{
				200,
				[]byte{'t', 'e', 's', 't'},
			},
			[]byte{'t', 'e', 's', 't'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResponse(tt.fields.StatusCode, tt.fields.Body)
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
