# Payments Processor

This application is a simplified payments processor.

---

## Pre-requiriments

`Docker Engine 19.03+`

`Docker-compose version 1.24+`

`Golang 1.13+`

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

*NOTE* : you must run `build` and start `dependencies` before.

```bash
make start-app
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

### `POST` /payment/cielo

Creates a new payment via Cielo. Eg.:

```bash
curl -v -X POST http://localhost:3000/payment/cielo -H 'Content-Type: application/json'  -d '{"order_id":"1b2c3d4e5f6g7h8i9j","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"authorized","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}'
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
curl -v -X POST http://localhost:3000/payment/rede -H 'Content-Type: application/json'  -d '{"order_id":"1b2c3d4e5f6g7h8i9j","customer":{"name":"lorem ipson"},"details":{"amount":100,"installments":1,"payment_type":"credit","card":{"source_id":"authorized","brand":"Visa","expiration_month":12,"expiration_year":2020},"itens":["lorem","ipson"]},"establishment":{"address":"rua lorem ipson","identifier":"00.111.222-8","postal_code":12345678}}'
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

*NOTE* : You can import the `postman collection` found at `docs/` folder to make easier to simulate all the errors.

---
