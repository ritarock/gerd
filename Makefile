.PHONY: build test

BINDIR=bin

build:
	go build -o $(BINDIR)/gerd ./main.go

test:
	go test ./...
