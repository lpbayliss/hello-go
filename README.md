# hello-go

Lightweight Go HTTP service with hello and health endpoints.

## Project Structure

```
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── config/          # Configuration loading
│   ├── greeting/        # Business logic
│   └── handler/         # HTTP handlers
├── build/               # Dockerfile
├── .golangci.yml        # Linter config
└── Makefile             # Build commands
```

## Endpoints

- `GET /hello` - Returns greeting message
  - `?name=John` - personalized greeting
  - `?uppercase=true` - uppercase output
- `GET /health` - Returns health status

## Development

```bash
make run           # run locally
make test          # run tests
make test-coverage # generate coverage report
make fmt           # format code
make lint          # run linter
make build         # build binary to bin/
make clean         # remove build artifacts
```

Server starts on port 8080 (override with `PORT` env var).

## Docker

```bash
make docker-build  # build image
make docker-run    # run container
```

Uses multi-stage build with scratch base (~5MB final image).

## Test

```bash
curl localhost:8080/hello
curl "localhost:8080/hello?name=John"
curl "localhost:8080/hello?name=John&uppercase=true"
curl localhost:8080/health
```

## Requirements

- Go 1.24+
- golangci-lint (for linting)
- goimports (for formatting)

### Install tools

```bash
brew install golangci-lint
go install golang.org/x/tools/cmd/goimports@latest
```
