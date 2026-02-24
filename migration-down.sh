docker exec -i duon_api migrate -path ./migrations -database "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable" down
