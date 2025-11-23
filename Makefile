.PHONY: build run test fmt lint clean docker-build docker-run

# Build the application
build:
	go build -o bin/server .

# Run the application locally
run:
	go run .

# Run all tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	gofmt -w .
	goimports -w .

# Run linter
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Build Docker image
docker-build:
	docker build -t hello-go .

# Run Docker container
docker-run:
	docker run -p 8080:8080 hello-go
