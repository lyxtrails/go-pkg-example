export GO111MODULE=on

SHELL := /bin/bash

build:
	go build ./...

clean:
	go clean -modcache

test:
	go test -v

cover:
	go test -v -covermode=count -coverprofile=count.out

cover-func:
	go tool cover -func=count.out

cover-html:
	go tool cover -html=count.out