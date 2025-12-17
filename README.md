# Backend Development Task

A RESTful API for managing users with name and date of birth, calculating age dynamically.

## Tech Stack

- Go
- Fiber (web framework)
- PostgreSQL 
- SQLC (code generation)
- Zap (logging)
- Validator (input validation)

## Setup

1. Install Go 1.21+
2. Install PostgreSQL
3. Clone the repo
4. Run `go mod tidy`
5. Create a `.env` file in the root directory:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=YOUR_POSTGRES_PASSWORD
   DB_NAME=backend_dev_task
   PORT=8080
   ```
6. Ensure PostgreSQL is running and the database exists
7. Run migrations: `go run cmd/server/main.go` (migrations run automatically)
8. Build: `go build ./cmd/server`
9. Run: `./server`

## Docker

1. `docker-compose up --build`

## API Endpoints

- POST /users - Create user
- GET /users/:id - Get user with age
- PUT /users/:id - Update user
- DELETE /users/:id - Delete user
- GET /users?limit=10&offset=0 - List users with pagination

## Testing

Run tests: `go test ./...`