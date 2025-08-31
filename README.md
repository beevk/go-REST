# REST API in Go

This project uses
- `go-chi/v5` for routing
- `golang-migrate/migrate` for database migrations


You can also install gin for hot-reloading during development:
```bash
go install github.com/codegangsta/gin@latest
``` 

## Setup
To install the dependencies, run:
```bash
go mod tidy
```
You will also need to run the migrations before starting the project:
```bash
make migrate-up
```
To rollback the last migration, run:
```bash
make migrate-down
```
To create a new migration, run:
```bash
make migrate-create name=<migration_name>
```

## Database
This project uses PostgreSQL as the database. You can use Docker to run a PostgreSQL container:
```bash
make docker-up
```
To stop the container, run:
```bash
make docker-down
```

## Starting the Project
To run the project, run the following commands:
Dev:
```bash
make run
```
Development with hot-reload:
```bash
make dev
```
Build:
```bash
make build
```
Prod:
```bash
make prod
```

