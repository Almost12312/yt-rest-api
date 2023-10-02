.PHONY: build
build:
	go build -v ./cmd/api-server
run:
	go run ./cmd/api-server
brun:
	go build -v ./cmd/api-server && ./api-server

.DEFAULT_GOAL:=brun