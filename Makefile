# Start Databases
infra:
	docker compose -f infra/docker-compose.infra.yml up -d

# Build and start everything
up:
	docker compose up --build

sqlc-auth:
	cd services/auth-service && sqlc generate

sqlc-exam:
	cd services/exam-service && sqlc generate
	
# Run Auth locally for debugging
auth:
	cd services/auth-service && go run cmd/main.go

# Run Exam locally
exam:
	cd services/exam-service && go run cmd/main.go

# Clean up everything
clean:
	docker compose down -v