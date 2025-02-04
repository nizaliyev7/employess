postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root employees
dropdb:
	docker exec -it postgres12  dropdb employees
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/employees?sslmode=disable" -verbose up	
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/employees?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server 