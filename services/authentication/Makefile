install:
	@go get -v
	@go install

run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f auth-service:0.0.1
	@docker-compose up -d

rm:
	@docker stop auth-service
	@docker rm -f auth-service
	@docker rmi -f auth-service:0.0.1