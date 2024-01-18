lint:
	@echo "running lint"
	@golangci-lint run ./...

external:
	@echo "run database using docker-compose"
	@docker network create app-default || true
	@docker compose --project-directory . -f deployments/docker-compose/docker-compose.yaml up --build

run:
	@go run cmd/application/main.go