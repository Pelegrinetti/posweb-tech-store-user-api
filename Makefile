build:
	@go build -o bin/server cmd/main.go

start-server:
	@go run cmd/main.go