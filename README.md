# An API Server boilerplate for Golang

Authored by [Marcus Ong](https://github.com/mraacus)

An easily extendible RESTful API server boilerplate in go.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Environment Setup](#environment-setup)
  - [Database Setup](#set-up-your-postgresql-database-with-docker)
  - [Running the Server](#running-the-server)
  - [Testing the API](#test-the-server-api)
- [Project Structure](#project-structure)
- [License](#license)

## Features

- Domain-driven design with modular architecture
- RESTful API setup built on Echo
- PostgreSQL integration with GORM for type-safe database access
- Database migrations with Goose for version-controlled schema changes
- A HTTP suite for type-safe request/response handling
- Validation using sonic and go-playground/validator
- Structured logging with slog
- Docker Integration for containerized PostgreSQL
- Hot reloading during development with air
- Environment configuration with godotenv


## Technologies

This repository uses the following technology stack.

- **[Echo](https://echo.labstack.com/)**: High performance, minimalist Go web framework
- **[GORM](https://gorm.io/)**: Industry standard ORM for Go
- **[Goose](https://github.com/pressly/goose)**: Database migration tool
- **[PostgreSQL](https://www.postgresql.org/)**: Robust, open-source relational database

## Prerequisites

- Go 1.25.0+
- Docker
- Goose
- Make (optional, for using Makefile commands)

## Getting Started

### Environment Setup

1. Clone the repository

   ```bash
   git clone https://github.com/mraacus/go-api-boilerplate-v2.git
   ```

2. Install dependencies
```bash
go mod tidy
```

3. Create your .env file
   ```bash
   cp .env.sample .env
   ```

### Set up your PostgreSQL database with Docker

This boilerplate comes with a docker-compose file to spin up a local PostgreSQL database.

1. Start the PostgreSQL container

   ```bash
   make docker-up
   # or
   docker-compose up -d
   ```

2. Run database migrations
   ```bash
   goose up
   ```

### Running the server

For development with hot reloading:

    make watch

Or run the server directly:

    go run cmd/api/main.go

### Test the server API

Testing just the API connection:

```
GET /

    Response:
    {
    "message": "I am groot"
    }

// Curl
curl -X GET http://localhost:8080/
```

Testing database functionality:

Create a user:

```

POST /users

    Request Body:
    {
    "name": "username",
    "role": "admin"
    }

    Response:
    {
    "id": 1,
    "name": "username",
    "role": "admin"
    }

// Curl
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "username",
    "role": "admin"
  }'
```

Get all users:

```
GET /users

    Response:
    [
      {
        "id": 1,
        "name": "username", 
        "role": "admin",
        "created_at": "2026-01-28T15:00:00Z",
        "updated_at": "2026-01-28T15:00:00Z"
      }
    ]

// Curl
curl -X GET http://localhost:8080/users
```

Ping your database to check its health using:

```
GET /health

curl -X GET http://localhost:8080/health
```

## Project Structure

```

.
├── cmd/
│ └── api/                  # Application entrypoint
├── handler/                # Echo HTTP handlers
├── internal/
│   ├── config/             # Configuration management
│   │   └── env.go          # Environment configuration
│   ├── dao/                # Data Access Objects (GORM models)
│   │   ├── user.go         # User model and database operations
│   │   └── ...
│   ├── db/                 # Database config and utilities
│   │   ├── migrations/     # Database migrations
│   │   ├── seed/           # Database seeding utilities
│   │   ├── db.go           # GORM database connection initialization
│   │   └── utils.go        # Database utilities
│   └── server/             # Server setup
│       ├── server.go       # Echo server initialization
│       └── routes.go       # Route definitions
├── pkg/                    # Shared packages
│   ├── common/             # Common utility packages
│   ├── constant/           # Constants
│   ├── domain/             # Domain model definitions
│   ├── external/           # External downstream integrations
│   ├── http/               # HTTP suite for request and response parsing
│   ├── middleware/         # Custom middlewares
│   └── validate/           # Custom validation
├── service/                # Service business logic
├── .env                    # Environment variables
├── docker-compose.yml      # Docker services
└── Makefile                # Make commands

```

## License

This boilerplate is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
