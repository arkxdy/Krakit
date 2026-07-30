# Contributing / local development

This guide covers how to run Krakit infrastructure and services with Docker.

## Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Compose v2)
- Go 1.23+ (only if you run services on the host)
- Make (optional; commands below also show raw `docker compose` equivalents)
- Copy env defaults:

```bash
cp .env.example .env
```

Edit `.env` if you need different ports or credentials. Docker Compose reads the root `.env` automatically.

## Layout

| Path | Purpose |
|------|---------|
| `docker-compose.yml` | Root compose; includes infra + services |
| `infra/docker-compose.infra.yml` | Postgres (auth/exam), MongoDB, Redis, Kafka, Kafka UI, MinIO |
| `infra/auth`, `infra/exam`, `infra/question` | DB init scripts |
| `services/docker-compose.services.yml` | `auth-service`, `exam-service` |
| `.env.example` | Documented env template |

## Quick start (full stack)

From the repo root:

```bash
cp .env.example .env
docker compose up --build -d
# or: make up
```

Check status:

```bash
docker compose ps
# or: make ps
```

Useful URLs once healthy:

| Service | URL |
|---------|-----|
| Auth API | http://localhost:8081/api/v1/health |
| Exam API | http://localhost:8082/api/v1/health |
| Kafka UI | http://localhost:8080 |
| MinIO Console | http://localhost:9001 (`minioadmin` / `minioadmin`) |

Stop (keep data):

```bash
docker compose down
# or: make down
```

Reset everything (deletes DB volumes):

```bash
docker compose down -v
# or: make clean
```

## Infra only (run Go services on the host)

Start dependencies:

```bash
docker compose up -d auth-db exam-db mongo-db auth-cache kafka kafka-ui minio
# or: make infra
```

Host ports (defaults from `.env.example`):

| Dependency | Host port |
|------------|-----------|
| Auth Postgres | `5432` |
| Exam Postgres | `5433` |
| MongoDB | `27017` |
| Redis | `6379` |
| Kafka (from host apps) | `29092` |
| Kafka (from other containers) | `kafka:9092` |
| MinIO API / Console | `9000` / `9001` |

Run services locally:

```bash
# terminal 1
make auth
# terminal 2 — exam defaults to exam DB on localhost:5433
make exam
```

Notes for host mode:

- Auth uses `DB_*` + `REDIS_*` from `.env` (defaults target auth DB on `5432`).
- Exam uses `DB_PORT=5433` and `DB_NAME=exam_db` by default; Kafka from the host must use `KAFKA_PORT=29092`.
- Exam JWKS default is `http://localhost:8081/api/v1/.well-known/jwks.json` — start auth first.

## Common commands

```bash
# Validate compose files (no containers started)
docker compose config

# Follow logs
docker compose logs -f
docker compose logs -f auth-service exam-service

# Rebuild a single service
docker compose up --build -d auth-service

# sqlc codegen
make sqlc-auth
make sqlc-exam
```

## Environment variables

See `.env.example` for the full list. Important groups:

- `AUTH_DB_*` / `EXAM_DB_*` — Postgres containers and credentials
- `MONGO_*` — Mongo root user + question DB name
- `REDIS_PASSWORD` — shared by Redis, auth-service, exam-service
- `MINIO_ACCESS_KEY` / `MINIO_SECRET_KEY` — object storage
- `AUTH_SERVICE_PORT` / `EXAM_SERVICE_PORT` — published app ports

When services run **inside** Compose, DB/Kafka/Redis hosts are container names (`auth-db`, `exam-db`, `kafka`, …). Those are set in `services/docker-compose.services.yml` and override host-oriented `.env` values for containers.

## Troubleshooting

**Postgres init scripts did not run**  
Init SQL only runs on first volume create. After changing `infra/auth` or `infra/exam`, reset volumes: `make clean` then `make infra`.

**Port already in use**  
Change the matching `*_PORT` in `.env` (for example `AUTH_DB_PORT` or `KAFKA_UI_PORT`).

**Exam cannot reach Mongo / Redis**  
Confirm infra is up (`docker compose ps`) and credentials match `.env`. Mongo requires `authSource=admin` (already built into exam-service config).

**Gateway missing**  
`gateway/` is not wired yet. Root compose only includes infra + auth/exam services. Add a gateway service to `docker-compose.yml` when that app exists.

## Suggested workflow

1. `cp .env.example .env`
2. `make infra` while iterating on Go code
3. `make auth` / `make exam` for fast local debugging
4. `make up` when you want the containerized services end-to-end
