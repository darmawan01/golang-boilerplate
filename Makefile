serve:
	@go run cmd/kodingworks/main.go

test:
	@go test -coverprofile=coverage.out ./...

migrate_clean:
	@rm db/migrations/*.sql

migrate:
	@cp -a */migrations/deploy/*.sql db/migrations
	@docker run --rm -v ${PWD}/db/migrations:/migrations --network host migrate/migrate\
    	-path=/migrations/ -database postgres://kodingworks:kodingworks@localhost:5432/kodingworks?sslmode=disable up

native_migrate:
	@cp -a */migrations/deploy/*.sql db/migrations
	@migrate -path db/migrations/ -database "postgres://kodingworks:kodingworks@localhost:5432/kodingworks?sslmode=disable" up

native_migrate_down:
	@migrate -path db/migrations/ -database "postgres://kodingworks:kodingworks@localhost:5432/kodingworks?sslmode=disable" down