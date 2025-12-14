# Octoback

A grocery list and recipe management app built with Go, Connect RPC, SQLite, and React.

## Prerequisites

- Go 1.22+
- Node.js 18+ and Yarn
- [sqlc](https://sqlc.dev/) for query generation
- [buf](https://buf.build/) for protobuf generation
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations

## Environment Variables

The following environment variables are required:

| Variable | Description |
|----------|-------------|
| `AUTH0_DOMAIN` | Auth0 domain (e.g., `your-tenant.us.auth0.com`) |
| `AUTH0_AUDIENCE` | Auth0 API audience identifier |

## Database Setup

Octoback uses SQLite as its database. The database file is created automatically when you run the server.

### Running Migrations

Before starting the server, apply the database migrations:

```bash
go run -tags 'sqlite' github.com/golang-migrate/migrate/v4/cmd/migrate -path=migrations -database="sqlite://./octoback.db" up
```

### Rolling Back Migrations

```bash
go run -tags 'sqlite' github.com/golang-migrate/migrate/v4/cmd/migrate -path=migrations -database="sqlite://./octoback.db" down
```

## Code Generation

Regenerate protobuf and SQL code:

```bash
go generate ./...
```

This runs:
- `sqlc generate` - generates Go code from SQL queries
- `buf generate` - generates Go and TypeScript code from protobuf definitions

## Running Locally

### 1. Start the Backend

```bash
go run ./cmd/apiserver -port 8080
```

The API server will be available at `http://localhost:8080`.

### 2. Start the Frontend

```bash
cd web
yarn install
yarn dev
```

The web app will be available at `http://localhost:5173` (or the next available port).

### Backend Configuration

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-port` | `PORT` | - | Port to listen on |
| `-db-path` | `OCTOMEAL_DB_PATH` | `octoback.db` | Path to SQLite database file |
| `-auth0-domain` | `AUTH0_DOMAIN` | - | Auth0 domain |
| `-auth0-audience` | `AUTH0_AUDIENCE` | - | Auth0 API audience |

## Running Tests

```bash
# Backend tests
go test ./...

# Frontend tests
cd web && yarn test
```

## API

The server exposes a Connect RPC API (gRPC-compatible) for managing grocery lists, recipes, and tenants.

### Services

- `GroceriesService` - CRUD operations for grocery lists and items
- `RecipesService` - CRUD operations for recipes and ingredients
- `TenantsService` - Tenant management and invitations

See `proto/octoback/` for the full API definitions.

