run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f user-service:0.0.1
	@docker compose up -d

rm:
	@docker stop user-service
	@docker rm -f user-service
	@docker rmi -f user-service:0.0.1