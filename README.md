# Octoback

A grocery list API built with Go, Connect RPC, and SQLite.

## Prerequisites

- Go 1.22+
- [sqlc](https://sqlc.dev/) for query generation
- [buf](https://buf.build/) for protobuf generation
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations

## Database Setup

Octoback uses SQLite as its database. The database file is created automatically when you run the server.

### Running Migrations

Before starting the server, apply the database migrations:

```bash
# Using golang-migrate
migrate -database "sqlite3://octoback.db" -path ./migrations up

# Or manually with sqlite3
sqlite3 octoback.db < migrations/000001_create_grocery_table.up.sql
sqlite3 octoback.db < migrations/000002_create_grocery_item_table.up.sql
```

### Rolling Back Migrations

```bash
migrate -database "sqlite3://octoback.db" -path ./migrations down
```

## Code Generation

Regenerate protobuf and SQL code:

```bash
go generate ./...
```

This runs:
- `sqlc generate` - generates Go code from SQL queries
- `buf generate` - generates Go code from protobuf definitions

## Running the Server

```bash
# With default settings (port from $PORT env, db at ./octoback.db)
go run ./cmd/apiserver

# With custom settings
go run ./cmd/apiserver -port 8080 -db-path ./data/groceries.db
```

### Configuration

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-port` | `PORT` | - | Port to listen on |
| `-db-path` | `OCTOMEAL_DB_PATH` | `octoback.db` | Path to SQLite database file |

## Running Tests

```bash
go test ./...
```

## API

The server exposes a Connect RPC API (gRPC-compatible) for managing grocery lists and items.

### Services

- `GroceryListService` - CRUD operations for grocery lists
- `GroceryItemService` - CRUD operations for grocery items

See `proto/octoback/groceries/v1/` for the full API definition.

