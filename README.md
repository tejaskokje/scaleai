# Go Project with Gin, File I/O, and Testing

This project demonstrates:
- HTTP server with Gin framework
- REST endpoints
- File operations (CSV and JSON)
- HTTP client for making requests
- Unit testing framework

## Installation

```
go mod tidy
```

## Running the server

```
go run main.go
```

## Testing the endpoint

```
curl http://localhost:8080/hello_world
```

## Running tests

```
go test -v
```

## Project Structure

- `main.go` - Entry point and Gin HTTP server
- `utils/` - Utility packages
  - `file_utils.go` - CSV and JSON file operations
  - `http_client.go` - HTTP client for making requests
- `main_test.go` - Unit tests
