createDb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=powtgres -d postgres:14-alpine

migrationup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

test:
	go test -v -cover ./...

.PHONY: createDb postgres  migrationup migrationdrop test