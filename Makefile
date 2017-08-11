setup:
	go get github.com/kelseyhightower/envconfig

run:
	go run main.go

fmt:
	@go fmt
	@cd commands & go fmt
	@cd configs & go fmt
	@cd handlers & go fmt
