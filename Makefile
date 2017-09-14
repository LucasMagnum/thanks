run:
	go run cmd/api.go

tests:
	go test -cover -v ./...

fmt:
	@go fmt ./cmd ./pkg

setup:
	go get ./...
