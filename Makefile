include scripts/*.mk

dev-build:
	docker compose -f docker-compose.dev.yaml build
dev-up: api_local_build dev-build
	docker compose -f docker-compose.dev.yaml up -d
dev-down:
	docker compose -f docker-compose.dev.yaml down
dev-logs:
	docker compose -f docker-compose.dev.yaml logs
dev-migrate-up:
	docker-compose -f docker-compose.dev.yaml up -d migrations-up
dev-migrate-down:
	docker compose --profile migrations-down -f docker-compose.dev.yaml up -d migrations-down
dev-api-run:
	go run cmd/api/api.go

