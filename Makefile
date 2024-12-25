postgres :
	docker run --name bank-postgres  --network banknet -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb :
	docker exec -it bank-postgres createdb --username=root --owner=root simplebank

dropdb :
	docker exec -it bank-postgres dropdb simplebank

migrateup :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose up
migrateup1step :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose up 1
migratedown :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose down

migratedown1step :
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose down 1

sqlc :
	sqlc generate

test :
	go test -v -cover ./...
server :
	go run main.go
mock:
	mockgen -package mockdb  -destination db/mock/store.go  github.com/phatvo2201/simplebank/db/sqlc Store

.PHONY: createdb postgres dropdb migratedown migrateup sqlc test server mock migratedown1step migrateup1step
