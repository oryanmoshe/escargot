.PHONY: build run

build:
	go build -o bin/escargot cmd/escargot/main.go

run:
	go run cmd/escargot/main.go
