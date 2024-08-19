BINARY_NAME=dns_server
SRC_DIR=./cmd/server
PKG_DIR=./pkg/...
GO_FILES=$(wildcard $(SRC_DIR)/*.go)

build:
	@echo "Building the project..."
	GO111MODULE=on go build -o $(BINARY_NAME) $(SRC_DIR)/main.go
	@chmod +x $(BINARY_NAME) # Ensure the binary is executable

run: build
	@echo "Running the server..."
	./$(BINARY_NAME)

serve: build
	@echo "Serving the project..."
	./$(BINARY_NAME)

dev: build
	@echo "Running in development mode..."
	# Install 'entr' if not already installed
	@if ! command -v entr &> /dev/null; then \
		brew install entr; \
	fi
	@find . -name '*.go' | entr -r make run

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

all: build

.PHONY: build run serve dev test clean all
