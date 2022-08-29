.PHONY: run test build fmt

run:
	go run main.go

test:
	go test -v

build:
	go build

fmt:
	go fmt ./...