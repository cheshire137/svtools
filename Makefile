all: build

default: build

build:
	go build -o bin/svtools svtools.go

test: build
	gotest ./...
