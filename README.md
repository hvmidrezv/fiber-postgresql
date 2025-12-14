# Fiber-Postgres Application

A professional REST API application built with Go Fiber framework and PostgreSQL database, following clean architecture principles.

## Features

- ✅ RESTful API with CRUD operations
- ✅ Clean Architecture with separation of concerns
- ✅ Repository Pattern for data access
- ✅ Dependency Injection
- ✅ Graceful shutdown
- ✅ Docker containerization
- ✅ Connection pooling
- ✅ Middleware (CORS, Logging, Recovery)
- ✅ Structured logging
- ✅ Error handling

## Project Structure

```
fiber-postgres/
├── config/                # Configuration management
├── handlers/             # HTTP request handlers
├── middleware/           # HTTP middleware
├── models/               # Data models and DTOs
├── repository/           # Data access layer
├── routes/               # Route definitions
├── storage/              # Database connections
└── main.go              # Application entry point
```

For detailed architecture documentation, see [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)

## Prerequisites

- Docker
- Docker Compose

## Quick Start

1. Start the application and database:
```bash
docker-compose up --build
```

2. The API will be available at `http://localhost:8080`

3. To run in detached mode:
```bash
docker-compose up -d --build
```

4. To stop the services:
```bash
docker-compose down
```

5. To stop and remove volumes:
```bash
docker-compose down -v
```

## API Endpoints

- `POST /api/create_books` - Create a new book
- `GET /api/get_books` - Get all books
- `GET /api/get_book/:id` - Get a book by ID
- `DELETE /api/delete_book/:id` - Delete a book by ID

## Example Request

Create a book:
```bash
curl -X POST http://localhost:8080/api/create_books \
  -H "Content-Type: application/json" \
  -d '{
    "author": "J.K. Rowling",
    "title": "Harry Potter",
    "publisher": "Bloomsbury"
  }'
```

Get all books:
```bash
curl http://localhost:8080/api/get_books
```

## Environment Variables

The application uses the following environment variables (defined in `.env` and `docker-compose.yml`):

- `DB_HOST` - PostgreSQL host (default: postgres)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - PostgreSQL user (default: postgres)
- `DB_PASSWORD` - PostgreSQL password (default: postgres123)
- `DB_NAME` - PostgreSQL database name (default: books_db)
- `DB_SSLMODE` - SSL mode (default: disable)

## Database

PostgreSQL database runs in a separate container with persistent storage using Docker volumes.

**Note:** PostgreSQL is exposed on port `5433` on the host (instead of the default 5432) to avoid conflicts with existing PostgreSQL installations. Inside the Docker network, it still uses port 5432.

## Development

To view logs:
```bash
docker-compose logs -f
```

To view logs for a specific service:
```bash
docker-compose logs -f app
docker-compose logs -f postgres
```

