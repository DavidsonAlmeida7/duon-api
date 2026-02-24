## Migrations
migrate-create:
	@read -p "Nome da migration: " name; \
	migrate create -ext sql -dir ./migrations -seq $$name

migrate-up:
	migrate -path ./migrations -database "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable" down 1