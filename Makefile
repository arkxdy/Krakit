.PHONY: help infra up down logs ps clean sqlc-auth sqlc-exam auth exam

help:
	@echo "Krakit Make targets"
	@echo "  make infra      Start infrastructure only (DBs, Redis, Kafka, MinIO)"
	@echo "  make up         Build and start full stack (infra + services)"
	@echo "  make down       Stop containers (keep volumes)"
	@echo "  make clean      Stop containers and remove volumes"
	@echo "  make logs       Tail compose logs"
	@echo "  make ps         Show running services"
	@echo "  make auth       Run auth-service on the host"
	@echo "  make exam       Run exam-service on the host"
	@echo "  make sqlc-auth  Generate auth-service sqlc code"
	@echo "  make sqlc-exam  Generate exam-service sqlc code"

# Start Databases / brokers / object storage
infra:
	docker compose up -d auth-db exam-db mongo-db auth-cache kafka kafka-ui minio

# Build and start everything
up:
	docker compose up --build -d

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

sqlc-auth:
	cd services/auth-service && sqlc generate

sqlc-exam:
	cd services/exam-service && sqlc generate

# Run Auth locally for debugging (expects make infra)
auth:
	cd services/auth-service && go run ./cmd/main.go

# Run Exam locally (expects make infra + auth reachable for JWKS)
exam:
	cd services/exam-service && go run ./cmd/main.go

# Clean up everything including volumes (destroys local DB data)
clean:
	docker compose down -v
