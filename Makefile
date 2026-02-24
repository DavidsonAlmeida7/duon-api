ifneq (,$(wildcard .env))
    include .env
    export
endif

SQL_FILE := internal/infra/database/db-fixtures/db-fixtures.sql

migrate-create:
	@read -p "Nome da migration: " name; \
	migrate create -ext sql -dir ./migrations -seq $$name

migrate-up:
	docker compose run --rm migrate -path=/migrations -database "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable" up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable" down 1

refresh-db:
	docker cp internal/infra/database/db-fixtures/db-fixtures.sql duon_db:/tmp/manual-seed.sql
	docker compose exec duon_db psql -U ${POSTGRES_USER} -d ${POSTGRES_DATABASE} -f /tmp/manual-seed.sql
	