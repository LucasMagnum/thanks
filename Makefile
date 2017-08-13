setup:
	go get github.com/kelseyhightower/envconfig

run:
	docker-compose run api go run main.go

integration-tests:
	go test -v tests/*

unit-tests:
	go test -v commands/*
	go test -v configs/*
	go test -v handlers/*


fmt:
	@go fmt
	@cd commands & go fmt
	@cd configs & go fmt
	@cd handlers & go fmt
