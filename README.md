# Blog API Go

# Description
An API for a blog post from my personal studying of gin and Go

## Stack
- Go
- Gin
- Postgres
- pgx
- migrate
- Docker compose

## Configuration
```env
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_PORT=5432
POSTGRES_HOST=localhost

PORT=3000
```

PORT is the port the API listen on

## How to run

1. Clone the repo

2. Run postgres (docker-compose)
  `docker compose up -d`

3. Migrate the database
  Migrate from the project root

  For fish users:
  ```fish
  bash -c 'set -a; source .env; set +a; migrate -path migrations -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" up'
  ```

  For bash and zsh users:
  ```bash
  set -a; source .env; set +a
  migrate -path migrations -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" up
  ```

4. Run the api:
  ```bash
  go run ./cmd/api/
  ```

## Endpoints

| Method | Route        | Description                                  |
|--------|--------------|----------------------------------------------|
| GET    | `/health`    | Checks if the API and the database are up    |
| POST   | `/posts`     | Creates a post                               |
| GET    | `/posts`     | Lists posts (optional filter: `?title=`)     |
| GET    | `/posts/:id` | Gets a post by id                            |
| PUT    | `/posts/:id` | Updates a post's title and content           |
| DELETE | `/posts/:id` | Deletes a post                               |


