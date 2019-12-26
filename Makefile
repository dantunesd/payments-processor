.PHONY: build start-app unit-test integration-test start-dep

APP=api
BIN=./bin/$(APP)

build:
	go build -o $(BIN) ./$(APP)

start-app:
	$(BIN)

unit-test:
	go test ./payment-processor

integration-test:
	go test ./integration-tests -count=1

start-dep:
	docker-compose up -d