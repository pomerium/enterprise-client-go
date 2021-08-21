export GO111MODULE=on

GO ?= "go"

.PHONY: test
test: 
	go test -v ./...
