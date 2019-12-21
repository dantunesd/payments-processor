.PHONY: build start-app unit-test start-dep

APP=api
BIN=./bin/$(APP)

build:
	go build -o $(BIN) ./$(APP)

start-app:
	$(BIN)

unit-test:
	go test ./...

start-dep:
	docker-compose up -d