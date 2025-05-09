run:
	go run cmd/main.go

migrate_up:
	goose -dir database/migrations postgres "postgres://postgres:penguin_dev@localhost:5432/shop" up
sqlc:
	sqlc generate --file=database/sqlc.yaml

docker_up:
	docker compose up -d

docker_down:
	docker compose down

.PHONY: run migrate_up sqlc docker_up docker_down