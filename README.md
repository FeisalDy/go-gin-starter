# Go Boilerplate

This is a boilerplate for a Go backend API using the Gin framework and PostgreSQL.

## Folder Structure

- `cmd/`: This directory contains the main application entry point. The `main.go` file in this directory is responsible for initializing the application, setting up the database connection, and starting the HTTP server.

- `internal/`: This directory contains all the private application code. This code is not meant to be imported by other applications.
  - `config/`: This directory contains the application configuration. The `config.go` file in this directory is responsible for loading the configuration from a `.env` file.
  - `database/`: This directory contains the database connection and migration logic. The `database.go` file in this directory is responsible for connecting to the database.
  - `models/`: This directory contains the data models for the application. Each file in this directory represents a single data model.
  - `handlers/`: This directory contains the HTTP handlers for the application. The handlers are responsible for handling incoming HTTP requests, calling the appropriate services, and returning an HTTP response.
  - `services/`: This directory contains the business logic for the application. The services are responsible for performing the core application logic, such as creating, reading, updating, and deleting data.
  - `repositories/`: This directory contains the database interaction logic for the application. The repositories are responsible for interacting with the database, such as creating, reading, updating, and deleting records.

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
```bash
go build -o my-app cmd/main.go
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
RUN go build -o my-app cmd/main.go

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