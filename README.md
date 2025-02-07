# Go REST API with Services & JWT

Go REST API using service principle, JWT auth, MySQL, and `httprouter`.

## Features

- RESTful API
- Service layer
- JWT authentication
- MySQL database
- `httprouter`

## Installation

1. `git clone [repo_url]`
2. `cd project_directory`
3. `go mod tidy`

## Configuration

Create `.env` and configure: `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_HOST`, `JWT_SECRET`.


## Authentication

JWT in `Authorization` header (`Bearer <token>`).

## Services

Business logic in service layer.

## Database

MySQL, schema in migrations (if applicable).


## Deployment

Standard Go deployment.
