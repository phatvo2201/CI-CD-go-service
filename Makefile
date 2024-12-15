postgres :
	docker run --name bank-postgres -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb :
	docker exec -it bank-postgres createdb --username=root --owner=root simplebank

dropdb :
	docker exec -it bank-postgres dropdb  simplebank
migrateup :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose up

migratedown :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose down
sqlc :
	slqc generate

.PHONY: createdb postgres dropdb migratedown migrateup sqlc