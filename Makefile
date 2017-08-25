run:
	docker-compose run --service-ports api go run cmd/api.go

test:
	go test -cover -v ./...

fmt:
	@go fmt ./...
