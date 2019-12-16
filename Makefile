.PHONY: run-build run-app run-unit-tests

run-build:
	go build -o ./bin/api ./api

run-app: 
	./bin/api

run-unit-tests:
	go test ./...

