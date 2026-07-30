# Krakit

Krakit — mock exam platform.

## Local development

See [CONTRIBUTING.md](./CONTRIBUTING.md) for Docker setup, Make targets, ports, and troubleshooting.

```bash
cp .env.example .env
make infra    # databases, Redis, Kafka, MinIO
make up       # full stack including auth-service + exam-service
```
