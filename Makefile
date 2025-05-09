run:
	go run cmd/main.go

migrate_up:
	goose -dir database/migrations postgres "postgres://postgres:penguin_dev@localhost:5432/shop" up
sqlc:
	sqlc generate --file=database/sqlc.yaml

.PHONY: run migrate_up sqlc