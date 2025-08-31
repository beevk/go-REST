.PHONY: run dev prod build docker-up docker-down

run: docker-up
	@echo "Starting server on port 8080..."
	@trap 'make docker-down' INT TERM; PORT=8080 go run main.go

dev: docker-up
	@echo "Starting development server with hot-reload on port 8080..."
	@trap 'make docker-down' INT TERM; PORT=8081 gin --port 8080 --appPort 8081 run main.go

prod:
	@echo "Starting production server on port 8080..."
	PORT=8080 go run main.go

build:
	@echo "Building binary..."
	go build -o app main.go

docker-up:
	@echo "Starting Docker Compose..."
	@./scripts/load-env.sh docker compose up -d

docker-down:
	@echo "Stopping Docker Compose..."
	docker compose down
