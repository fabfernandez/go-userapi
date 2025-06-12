# User API

A simple CRUD REST API microservice for managing users, built with Go and MySQL.

This whole thing was made in a few hours in 1 day using the Cursor IDE, using Claude 3.5 as the IA model.
It's a fully functional docker app. Includes tests, documentation on how to start it and how to use the endpoints.
I have no idea about Go, I neved coded in Go before and this doesn't count.
This works as an example about how the software development profession is changing.

## Features

- CRUD operations for users
- Input validation
- Static API documentation with Swagger UI
- Docker support
- MySQL database
- Built using Test-driven development (I insisted a lot in my prompts)
- Makefile for easy development

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- MySQL (if running locally)
- Make

## Getting Started

### Using Make (Recommended)

The project includes a Makefile with common operations. To see all available commands:
```bash
make help
```

1. Clone the repository:
```bash
git clone <repository-url>
cd userapi
```

2. Initialize the database:
```bash
make init-db
```

3. Run the application:
```bash
make run
```

Or with custom configuration:
```bash
DB_USER=myuser DB_PASSWORD=mypassword make run
```

### Using Docker

1. Start the application:
```bash
make docker-up
```

2. Stop the application:
```bash
make docker-down
```

The API will be available at `http://localhost:8080`
The API documentation will be available at `http://localhost:8080/docs/`

### Manual Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd userapi
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the database:
```bash
mysql -u root -p < schema.sql
```

4. Set environment variables (or use defaults):
```bash
export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=root
export DB_NAME=userdb
export DB_PORT=3306
export PORT=8080
```

5. Run the application:
```bash
go run main.go
```

## Development

### Running Tests
```bash
make test
```

### Building the Application
```bash
make build
```

### Cleaning Build Artifacts
```bash
make clean
```

## API Documentation

API documentation is available at `http://localhost:8080/docs/`

### Endpoints

- `POST /users` - Create a new user
- `GET /users/{id}` - Get a user by ID
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user
- `GET /users` - List all users

