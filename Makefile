.PHONY: build start unit-test integration-test start-dep stop-dep

APP=api
BIN=./bin/$(APP)

build:
	go build -o $(BIN) ./$(APP)

start:
	$(BIN)

unit-test:
	go test ./payment-processor -count=1

integration-test:
	go test ./integration-tests -count=1

start-dep:
	docker-compose up

stop-dep:
	docker-compose stop