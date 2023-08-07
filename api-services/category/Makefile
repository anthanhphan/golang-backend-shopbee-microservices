run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f category-service:0.0.1
	@docker compose up -d

rm:
	@docker stop category-service
	@docker rm -f category-service
	@docker rmi -f category-service:0.0.1