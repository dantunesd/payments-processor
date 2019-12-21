package payment

import (
	"errors"
	"reflect"
	"testing"
)

func TestCieloRepository_Sale(t *testing.T) {
	type fields struct {
		r IHTTPRequester
	}
	type args struct {
		crb CieloRequestBody
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CieloResponseBody
		wantErr bool
	}{
		{
			"request with success",
			fields{
				r: requesterMock{
					post: func(path string, body, output interface{}) error {
						return nil
					},
				},
			},
			args{
				CieloRequestBody{},
			},
			&CieloResponseBody{},
			false,
		},
		{
			"fail to request",
			fields{
				r: requesterMock{
					post: func(path string, body, output interface{}) error {
						return errors.New("failed for some reason")
					},
				},
			},
			args{
				CieloRequestBody{},
			},
			&CieloResponseBody{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCieloRepository(tt.fields.r)
			got, err := c.Sale(tt.args.crb)
			if (err != nil) != tt.wantErr {
				t.Errorf("CieloRepository.Sale() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CieloRepository.Sale() = %v, want %v", got, tt.want)
			}
		})
	}
}

type requesterMock struct {
	post func(path string, body, output interface{}) error
}

func (r requesterMock) Post(path string, body, output interface{}) error {
	return r.post(path, body, output)
}
