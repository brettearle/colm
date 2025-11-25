# Binary name
BINARY := colm-broker

# Command path
CMD := ./cmd/colm-broker

# Default target
.PHONY: build
build:
	go build -o bin/$(BINARY) $(CMD)

.PHONY: run
run:
	go run $(CMD)

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	go fmt ./...
