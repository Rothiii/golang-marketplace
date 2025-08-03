build:
	@go build -o bin/golang-marketplace cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/golang-marketplace