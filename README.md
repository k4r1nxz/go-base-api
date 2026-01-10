# Go Base API

A foundational API template built with Go, featuring routing, middleware, configuration management, and CRUD operations.

## Features

- RESTful API endpoints
- Middleware for logging and CORS
- Configuration management with environment variables
- Basic CRUD operations for users
- Structured project layout

## Prerequisites

- Go 1.19 or higher

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd go-base-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables (optional, see [Configuration](#configuration)):
   ```bash
   export PORT=8080
   export DB_HOST=localhost
   export DB_PORT=5432
   # ... other variables
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## Configuration

The application uses environment variables for configuration. Default values are:

- `PORT`: Port to run the server on (default: "8080")
- `DB_HOST`: Database host (default: "localhost")
- `DB_PORT`: Database port (default: "5432")
- `DB_USER`: Database user (default: "user")
- `DB_PASSWORD`: Database password (default: "password")
- `DB_NAME`: Database name (default: "mydb")
- `JWT_SECRET`: JWT secret key (default: "default_secret_key")

## API Endpoints

- `GET /` - Home endpoint
- `GET /health` - Health check
- `GET /api/users` - Get all users
- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user by ID
- `DELETE /api/users/{id}` - Delete user by ID

## Project Structure

```
go-base-api/
├── main.go          # Entry point of the application
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
├── config/          # Configuration management
│   └── config.go
├── handlers/        # Request handlers
│   └── handlers.go
├── middleware/      # Middleware functions
│   └── middleware.go
├── models/          # Data models (currently empty)
└── utils/           # Utility functions (currently empty)
```

## Example Requests

### Get all users
```bash
curl http://localhost:8080/api/users
```

### Create a user
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"New User","email":"newuser@example.com"}'
```

### Get user by ID
```bash
curl http://localhost:8080/api/users/1
```

## Dependencies

This project uses the following external packages:
- `github.com/gorilla/mux` - HTTP request multiplexer

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a pull request

## License

This project is licensed under the MIT License.