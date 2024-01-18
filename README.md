# Pilotoapi
Pilotoapi is a simple API written in [Go](#Golang) to manage products from [MySQL](#MySQL) database.

## Setup
- Use `make` to build and run
- [Go 1.21](#Golang)
- [Docker](#Docker)
- [Docker Compose](#Docker)
- [MySQL 5.7](#MySQL)
- [Postman](#Postman)

## Usage
```bash
make external
make run
```

## Endpoints
```bash
### GET /api/product
### GET /api/product/{id}
### POST /api/product
### PUT /api/product/{id}
### DELETE /api/product/{id}
```
## Postman
- In the root of the project, there is postman collection that can be import to your own environment to use routes
