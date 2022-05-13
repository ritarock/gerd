.PHONY: install build test

BINDIR=bin
INSTALL_PATH=github.com/ritarock/gerd

install:
	go install $(INSTALL_PATH)

build:
	go build -o $(BINDIR)/gerd ./main.go

test:
	go test $(shell go list ./... | grep internal)
