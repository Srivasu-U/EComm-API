build:
	@go build -o bin/EComm-API cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/EComm-API