# User API

A simple REST API for managing users, built with Go and MySQL.

## Features

- CRUD operations for users
- Input validation
- Static API documentation with Swagger UI
- Docker support
- MySQL database
- Test-driven development
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

### Example Request

Create a new user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "age": 30,
    "phone_number": "+1234567890",
    "email": "john@example.com"
  }'
```

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
├── handlers/
│   ├── user_handler.go
│   └── user_handler_test.go
├── main.go
├── models/
│   ├── user.go
│   └── user_test.go
├── repository/
│   ├── mysql_user_repository.go
│   └── user_repository.go
├── docs/
│   ├── swagger.json
│   └── index.html
└── schema.sql
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 