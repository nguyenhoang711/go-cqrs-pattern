run:
	go run cmd/main.go

migrate-up:
	goose -dir database/migrations postgres "postgres://postgres:penguin_dev@localhost:5432/shop" up
sqlc:
	sqlc generate --file=database/sqlc.yaml

docker-up:
	docker compose up -d

docker-down:
	docker compose down

generate-proto:
	cd api && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. order.proto
	
.PHONY: run migrate-up sqlc docker-up docker-down generate-proto