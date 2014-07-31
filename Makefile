all: build

deps:
	go get -d

build: deps
	go build

test: deps
	go test ./...

install: deps
	go install

