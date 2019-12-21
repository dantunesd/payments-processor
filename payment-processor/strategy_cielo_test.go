package payment

import (
	"errors"
	"testing"
)

func TestCieloStrategy_Process(t *testing.T) {
	type fields struct {
		r ICieloRepository
	}
	type args struct {
		p Payment
		s Source
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"process with success",
			fields{
				r: repositoryMock{
					func(CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{}, nil
					},
				},
			},
			args{
				Payment{},
				Source{},
			},
			false,
		},
		{
			"failed to process",
			fields{
				r: repositoryMock{
					func(CieloRequestBody) (*CieloResponseBody, error) {
						return &CieloResponseBody{}, errors.New("failed to request")
					},
				},
			},
			args{
				Payment{},
				Source{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CieloStrategy{r: tt.fields.r}
			if err := c.Process(tt.args.p, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CieloStrategy.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type repositoryMock struct {
	sale func(CieloRequestBody) (*CieloResponseBody, error)
}

func (r repositoryMock) Sale(c CieloRequestBody) (*CieloResponseBody, error) {
	return r.sale(c)
}
