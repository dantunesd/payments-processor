package payment

import (
	"reflect"
	"testing"
)

func TestNewError(t *testing.T) {
	type args struct {
		message    string
		errType    ErrorType
		statusCode int
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
				InvalidContentError,
				400,
			},
			&Error{
				ErrorMessage: "error",
				ErrorType:    InvalidContentError,
				StatusCode:   400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.message, tt.args.errType, tt.args.statusCode); !reflect.DeepEqual(got, tt.want) {
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
				ErrorType:    InvalidContentError,
				StatusCode:   400,
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

func TestNewInternalServerError(t *testing.T) {
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
				ErrorType:    InternalServerError,
				StatusCode:   500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalServerError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransactionError(t *testing.T) {
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
				ErrorType:    TransactionError,
				StatusCode:   400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionError() = %v, want %v", got, tt.want)
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
