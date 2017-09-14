run:
	go run cmd/api.go

tests:
	go test -cover -v ./...

fmt:
	@go fmt ./...

setup:
	go get ./...
