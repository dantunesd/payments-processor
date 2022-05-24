# Payments Processor

This application is a simplified payments processor.

---

## Pre-requiriments

The application was developed with the dependencies and versions bellow:

`Docker Engine - 18.06.0+`

`Docker-compose - 1.24.1+`

`Golang - 1.13+`

--- 

## Build

Run the command bellow to generate the executable:

```bash
make build
```

---

## Dependencies

Run the command bellow to start the dependencies:

```bash
make start-dep
```

---

## Environment

Enviroment variables are parsed by `api/config.go` file. It contains all the default values for development purposes.

To change some variable value, you must sets it in your environment. Eg. :

```bash
export APP_NAME=payment-processor
```

---

## Application

Run the command bellow to start the application:

*NOTE* : you must run `build` and start `dependencies` before. Wait until database service be ready. 

```bash
make start
```

---

## Unit Tests

Run the command bellow to run all unit tests:

```bash
make unit-test
```

---

## Integration Tests

Run the command bellow to run all integration tests:

*NOTE* : you must start `dependencies` and start `application` before.

```bash
make integration-test
```

---

## API

*NOTE* : You can donwload and import the [postman collection](https://github.com/dantunesd/payments-processor/blob/master/docs/payments-processor.postman_collection.json) to make easier to do the requests below.

### `POST` /payment/cielo

Creates a new payment via Cielo. Eg.:

```bash
curl -v -X POST http://localhost:3000/payment/cielo -H 'Content-Type: application/json'  -d '{"order_id":"1b2c3d4e5f6g7h8i9j","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}'
```

### Responses

| http status code | description |
| :---------- | :--------------- |
| `200` | {"message": "payment succeeded"}|
| `400` | {"message": "some friendly message", "error": "the root cause"}|
| `500` | {"message": "some friendly message", "error": "the root cause"}|


### `POST` /payment/rede

Creates a new payment via Rede. Eg.:

```bash
curl -v -X POST http://localhost:3000/payment/rede -H 'Content-Type: application/json'  -d '{"order_id":"1b2c3d4e5f6g7h8i9j","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}'
```

### Responses

| http status code | description |
| :---------- | :--------------- |
| `200` | {"message": "payment succeeded"}|
| `400` | {"message": "some friendly message", "error": "the root cause"}|
| `500` | {"message": "some friendly message", "error": "the root cause"}|

---

## Development

In development environment (running the application and dependencies locally), you can change `order_id` field value in the request body on `/payment/*acquirer*` to see the application behaviors. Eg.:

| order_id value | with acquirer |
| :---------- | :--------------- |
| integration-error | cielo/rede |
| emissor-error | cielo/rede |
| business-error | only on rede |


*NOTE* : You can donwload and import the [postman collection](https://github.com/dantunesd/payments-processor/blob/master/docs/payments-processor.postman_collection.json) to make easier to simulate all the errors.

---

## Sandbox

To test with acquirers sandbox, you must set the following environment variables with acquirer's sandbox values:

```bash
export REDE_URI="https://api.userede.com.br/desenvolvedores"
export REDE_AUTH="Basic your-basic-auth"
export CIELO_URI="https://apisandbox.cieloecommerce.cielo.com.br"
export CIELO_MERCHANT_ID="your-merchant-id"
export CIELO_MERCHANT_KEY="your-merchant-key"
```

After that you must use the inserteds cards (`source_id`) at database (see `docker-data/database/db.sql`) for cielo and rede on the request payload.

You can try to change some values (like `amount`) on the payload to receive other reponses.

See Rede/Cielo sandbox documentation to see how to get new responses:

https://developercielo.github.io/manual/cielo-ecommerce#cart%C3%A3o-de-cr%C3%A9dito-sandbox

https://www.userede.com.br/desenvolvedores/pt/produto/e-Rede#tutorial-playground

https://www.userede.com.br/desenvolvedores/pt/produto/e-Rede#tutorial-erros

---
