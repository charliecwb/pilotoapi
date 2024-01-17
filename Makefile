DATABASE_NAME=app
DB_USR=root
DB_PWD=root
DB_HOST=localhost
DB_PORT=3306

lint:
	@echo "running lint"
	@golangci-lint run ./...

external:
	@echo "run database using docker-compose"
	@docker network create app-default || true
	@docker compose --project-directory . -f deployments/docker-compose/docker-compose.yaml up --build

build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/app github.com/pilotoAPI/cmd/application/

run:
	@go run cmd/app/main.go