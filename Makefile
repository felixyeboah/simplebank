postgres:
	docker run --name postgres16 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres16 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server