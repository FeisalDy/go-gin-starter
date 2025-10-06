# Go Boilerplate

This is a boilerplate for a Go backend API using the Gin framework and PostgreSQL.

## Folder Structure

- `cmd/server/`: This directory contains the main application entry point. The `main.go` file in this directory is responsible for initializing the application, setting up the database connection, and starting the HTTP server.

- `config/`: This directory contains the application configuration. The `config.go` file in this directory is responsible for loading the configuration from a `.env` file.

- `internal/`: This directory contains all the private application code, organized by domain.
  - `user/`: This directory contains all the code related to the user domain.
    - `handler/`: This directory contains the HTTP handlers for the user domain.
    - `service/`: This directory contains the business logic for the user domain.
    - `repository/`: This directory contains the database interaction logic for the user domain.
    - `model/`: This directory contains the data models for the user domain.
    - `dto/`: This directory contains the data transfer objects for the user domain.
  - `common/`: This directory contains shared utilities, middleware, and error handling code.
    - `middleware/`: This directory contains custom middleware for the application.
    - `errors/`: This directory contains custom error types for the application.
    - `utils/`: This directory contains utility functions for the application.

- `pkg/`: This directory contains any public-facing libraries or utilities that can be imported by other applications.

- `scripts/`: This directory contains any helper scripts for the application, such as a script to run the application.

## Development

### Prerequisites

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Air](https://github.com/cosmtrek/air)

### Setup

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Create a `.env` file in the root directory with the following variables:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=postgres
    ```
4.  Install `air` for hot reloading:
    ```bash
    go install github.com/cosmtrek/air@latest
    ```
5.  Run the application with `air`:
    ```bash
    air
    ```
    This will start the application and automatically reload it when you make changes to the code.

## Building and Deployment

### Building

To build the application, run the following command:
go build -o my-app cmd/server/main.go
```
This will create a binary file named `my-app` in the root directory.

### Deployment

To deploy the application, you can simply run the binary file:
```bash
./my-app
```

You can also use a process manager like `systemd` or `supervisor` to run the application in the background.

For a more robust deployment, you can use Docker to containerize the application. Here is an example `Dockerfile`:

```Dockerfile
# Start from the official Go image
FROM golang:1.21-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o my-app cmd/server/main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./my-app"]
```

You can then build the Docker image and run the container:

```bash
docker build -t my-app .
docker run -p 8080:8080 my-app
```