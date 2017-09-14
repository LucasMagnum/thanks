run:
	go run cmd/api.go

tests:
	go test -cover -v ./app

fmt:
	@go fmt ./ ./app

setup:
	go get ./...
