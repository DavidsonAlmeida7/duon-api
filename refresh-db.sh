docker exec -i emcash_db mariadb -u root -p$1 -D dbapp < db-fixtures/pre-fixtures.sql

docker exec -i emcash_simulador migrate -path ./infra/database/migrations/mariaDB -database "mysql://root:$1@tcp(emcash_db)/dbsimulador" down

docker exec -i emcash_simulador migrate -path ./infra/database/migrations/mariaDB -database "mysql://root:$1@tcp(emcash_db)/dbsimulador" up

docker exec -i emcash_db mariadb -u root -p$1 -D dbsimulador < db-fixtures/fixtures.sql
