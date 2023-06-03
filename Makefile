install:
	@go get -v
	@go install

dev:
	@cd cmd/dev;go run main.go

pro:
	@cd cmd/pro;go run main.go

docker-run:
	@docker rmi -f shopbee-api:1.0
	@docker-compose up