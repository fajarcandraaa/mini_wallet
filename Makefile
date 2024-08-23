#!/bin/bash
update:
	@go get -u
	
tidy:
	@go mod tidy
	@go mod vendor

clean:
	@go clean -cache
	
test-service:
	@go test  -cover ./internal/service

test-cover:
	@go test $$(go list ./internal/service  | grep -v /vendor/) -coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.out

start:
	@go run main.go

