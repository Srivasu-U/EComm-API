build:
	@go build -o bin/EComm-API cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/EComm-API

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down