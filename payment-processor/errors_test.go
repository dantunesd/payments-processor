package payment

import (
	"reflect"
	"testing"
)

func TestNewError(t *testing.T) {
	type args struct {
		message string
		errType ErrorType
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			"return a new error",
			args{
				"error",
				InvalidContent,
			},
			&Error{
				ErrorMessage: "error",
				ErrorType:    InvalidContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.message, tt.args.errType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidContentError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			"return a error",
			args{
				message: "error",
			},
			&Error{
				ErrorMessage: "error",
				ErrorType:    InvalidContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvalidContentError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvalidContentError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidRequestError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			"return a error",
			args{
				message: "error",
			},
			&Error{
				ErrorMessage: "error",
				ErrorType:    InvalidRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvalidRequestError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvalidRequestError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPaymentFailedError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			"return a error",
			args{
				message: "error",
			},
			&Error{
				ErrorMessage: "error",
				ErrorType:    PaymentFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentFailedError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentFailedError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		ErrorMessage string
		ErrorType    ErrorType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"return a message",
			fields{
				ErrorMessage: "error",
			},
			"error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				ErrorMessage: tt.fields.ErrorMessage,
				ErrorType:    tt.fields.ErrorType,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
