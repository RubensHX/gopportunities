.PHONY: default run build test docs clean

APP_NAME = "gopportunities"

default: run-with-docs

run:
	@echo "Running $(APP_NAME)..."
	@go run main.go
run-with-docs:
	@echo "Running $(APP_NAME) with docs..."
	@swag init
	@go run main.go -docs
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) main.go
test:
	@echo "Testing $(APP_NAME)..."
	@go test ./...
docs:
	@echo "Generating docs for $(APP_NAME)..."
	@swag init
clean:
	@echo "Cleaning $(APP_NAME)..."
	@rm -rf $(APP_NAME) ./docs