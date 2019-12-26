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