install:
	@go get -v
	@go install

dev:
	@cd cmd/dev;go run main.go

pro:
	@cd cmd/pro;go run main.go