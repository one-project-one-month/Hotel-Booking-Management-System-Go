.PHONY: dev
dev:
	@echo "Starting development server..."
	air -c .air.toml

.PHONY: build
build:
	@echo "Building the application..."
	go build -o ./bin/app ./cmd/app

.PHONY: run
run:
	@echo "Running the application..."
	go run ./cmd/app/main.go

.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

.PHONE: format
format:
	@echo "Formating the code..."
	go fmt ./...

.PHONY: lint
lint:
	@echo "Running linters..."
	golangci-lint run

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf ./bin
	rm -rf ./tmp

.PHONY: install
install:
	@echo "Installing dependencies..."
	go install github.com/air-verse/air@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
	go mod tidy

all: clean install lint test build