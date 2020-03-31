serve:
	@LOCAL=true go run cmd/kodingworks/main.go

test:
	@go test -coverprofile=coverage.out ./tests/...

migrate_clean:
	@rm db/migrations/*.sql

migrate:
	@cp -a */migrations/deploy/*.sql db/migrations
	@docker run --rm -v ${PWD}/db/migrations:/migrations --network host migrate/migrate\
    	-path=/migrations/ -database postgres://kodingworks:kodingworks@localhost:65432/kodingworks?sslmode=disable up

native_migrate:
	@cp -a */migrations/deploy/*.sql db/migrations
	@migrate -path db/migrations/ -database "postgres://kodingworks:kodingworks@localhost:65432/kodingworks?sslmode=disable" up

heroku_migrate:
	@cp -a */migrations/deploy/*.sql db/migrations
	@migrate -path db/migrations/ -database "postgres://obqutqnwosixcv:7ffa0b47a51d0104a57731a2a5c4b3043cc5834b2c1d5d3f6b6c30a43f03d442@ec2-54-157-78-113.compute-1.amazonaws.com:5432/d3ll9fit5qi1bh" up