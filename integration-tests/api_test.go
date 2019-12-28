package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		method      string
		body        []byte
		wantStatus  int
		wantReponse string
	}{
		{
			"Process a cielo payment with success",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{"order_id":"1b2c3d4","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			200,
			`{"message":"payment succeeded"}`,
		},
		{
			"Process a cielo payment with inexistent source id",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{"order_id":"1b2c3d4","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"inexistent","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"source_id not found"}`,
		},
		{
			"Process a cielo payment with a integration error",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{"order_id":"integration-error","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"ReturnCode: 177, ReturnMessage: Card Number is invalid"}`,
		},
		{
			"Process a cielo payment with a emissor error",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{"order_id":"emissor-error","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"ReturnCode: 05, ReturnMessage: Not Authorized"}`,
		},
		{
			"Process a cielo payment with required fields missing",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{"customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"Key: 'Payment.OrderID' Error:Field validation for 'OrderID' failed on the 'required' tag"}`,
		},
		{
			"Process a cielo payment without all fields",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`{}`),
			400,
			`{"message":"Failed to proccess payment","error":"Key: 'Payment.OrderID' Error:Field validation for 'OrderID' failed on the 'required' tag\nKey: 'Payment.Customer.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Payment.Details.Card.SourceID' Error:Field validation for 'SourceID' failed on the 'alphanum' tag\nKey: 'Payment.Details.Card.Brand' Error:Field validation for 'Brand' failed on the 'alpha' tag\nKey: 'Payment.Details.Card.ExpirationYear' Error:Field validation for 'ExpirationYear' failed on the 'required' tag\nKey: 'Payment.Details.Card.ExpirationMonth' Error:Field validation for 'ExpirationMonth' failed on the 'min' tag\nKey: 'Payment.Details.Amount' Error:Field validation for 'Amount' failed on the 'min' tag\nKey: 'Payment.Details.PaymentType' Error:Field validation for 'PaymentType' failed on the 'alpha' tag\nKey: 'Payment.Details.Installments' Error:Field validation for 'Installments' failed on the 'min' tag\nKey: 'Payment.Details.Itens' Error:Field validation for 'Itens' failed on the 'gte' tag\nKey: 'Payment.Establishment.Identifier' Error:Field validation for 'Identifier' failed on the 'required' tag\nKey: 'Payment.Establishment.Address' Error:Field validation for 'Address' failed on the 'required' tag\nKey: 'Payment.Establishment.PostalCode' Error:Field validation for 'PostalCode' failed on the 'required' tag"}`,
		},
		{
			"Process a cielo payment with invalid content",
			"http://localhost:3000/payment/cielo",
			http.MethodPost,
			[]byte(`a`),
			400,
			`{"message":"Invalid body content","error":"invalid character 'a' looking for beginning of value"}`,
		},
		{
			"Process a rede payment with success",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"order_id":"1b2c3d4","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			200,
			`{"message":"payment succeeded"}`,
		},
		{
			"Process a rede payment with inexistent source id",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"order_id":"1b2c3d4","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"inexistent","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"source_id not found"}`,
		},
		{
			"Process a rede payment with a integration error",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"order_id":"integration-error","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"ReturnCode: 37, ReturnMessage: CardNumber: Invalid parameter format."}`,
		},
		{
			"Process a rede payment with a emissor error",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"order_id":"emissor-error","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"ReturnCode: 101, ReturnMessage: Unauthorized. Problems on the card, contact the issuer."}`,
		},
		{
			"Process a rede payment with a business error",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"order_id":"business-error","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"ReturnCode: 42, ReturnMessage: Reference: Order number already exists."}`,
		},
		{
			"Process a rede payment with required fields missing",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{"customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}`),
			400,
			`{"message":"Failed to proccess payment","error":"Key: 'Payment.OrderID' Error:Field validation for 'OrderID' failed on the 'required' tag"}`,
		},
		{
			"Process a rede payment without all fields",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`{}`),
			400,
			`{"message":"Failed to proccess payment","error":"Key: 'Payment.OrderID' Error:Field validation for 'OrderID' failed on the 'required' tag\nKey: 'Payment.Customer.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Payment.Details.Card.SourceID' Error:Field validation for 'SourceID' failed on the 'alphanum' tag\nKey: 'Payment.Details.Card.Brand' Error:Field validation for 'Brand' failed on the 'alpha' tag\nKey: 'Payment.Details.Card.ExpirationYear' Error:Field validation for 'ExpirationYear' failed on the 'required' tag\nKey: 'Payment.Details.Card.ExpirationMonth' Error:Field validation for 'ExpirationMonth' failed on the 'min' tag\nKey: 'Payment.Details.Amount' Error:Field validation for 'Amount' failed on the 'min' tag\nKey: 'Payment.Details.PaymentType' Error:Field validation for 'PaymentType' failed on the 'alpha' tag\nKey: 'Payment.Details.Installments' Error:Field validation for 'Installments' failed on the 'min' tag\nKey: 'Payment.Details.Itens' Error:Field validation for 'Itens' failed on the 'gte' tag\nKey: 'Payment.Establishment.Identifier' Error:Field validation for 'Identifier' failed on the 'required' tag\nKey: 'Payment.Establishment.Address' Error:Field validation for 'Address' failed on the 'required' tag\nKey: 'Payment.Establishment.PostalCode' Error:Field validation for 'PostalCode' failed on the 'required' tag"}`,
		},
		{
			"Process a rede payment with invalid content",
			"http://localhost:3000/payment/rede",
			http.MethodPost,
			[]byte(`a`),
			400,
			`{"message":"Invalid body content","error":"invalid character 'a' looking for beginning of value"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var content io.Reader
			if len(tt.body) > 0 {
				content = bytes.NewBuffer(tt.body)
			}

			req, _ := http.NewRequest(tt.method, tt.url, content)

			h := http.Client{}
			resp, err := h.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if status := resp.StatusCode; status != tt.wantStatus {
				t.Errorf("status code: got %v want %v",
					status, tt.wantStatus)
			}

			r, _ := ioutil.ReadAll(resp.Body)
			if string(r) != tt.wantReponse {
				t.Errorf("body: got %v want %v",
					string(r), tt.wantReponse)
			}
		})
	}
}
