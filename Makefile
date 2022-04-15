.PHONY: default usage build run build-release

usage:
	@echo "Please provide an option:"
	@echo "	make build		--- Build the app"
	@echo "	make run		--- Run the app"
	@echo "	make build-release	--- Build with optimization flags enabled"

build:
	go build -o simpleserver.out cmd/simpleserver/main.go

run:
	go run cmd/simpleserver/main.go

build-release:
	go build -o simpleserver.out -ldflags "-s -w" cmd/simpleserver/main.go

default: usage
