User Management REST API

A RESTful API built using Go (Fiber) to manage users with name and date of birth.
The API calculates the user’s age dynamically at runtime using Go’s `time` package
and does not store age in the database.

Features
User CRUD operations
Date of birth (`dob`) stored in the database
Age calculated dynamically using Go’s `time` package
Input validation with go-playground/validator
SQLC for type-safe database access
Structured logging using Uber Zap
Request ID injected into response headers
Pagination support for listing users
MySQL database with Docker support

Tech Stack
Go
Fiber
MySQL
SQLC
Docker
Zap

API Endpoints

| Method | Endpoint        | Description          |
|------|-----------------|----------------------|
| POST | `/users`        | Create user          |
| GET  | `/users`        | List users           |
| GET  | `/users/:id`    | Get user by ID       |
| PUT  | `/users/:id`    | Update user          |
| DELETE | `/users/:id`  | Delete user          |

Running the Project

Start database
```bash
docker-compose up -d

start server
go run ./cmd/server