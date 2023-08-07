run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f mail-service:0.0.1
	@docker compose up -d

rm:
	@docker stop mail-service
	@docker rm -f mail-service
	@docker rmi -f mail9-service:0.0.1