run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f cart-service:0.0.1
	@docker compose up -d

rm:
	@docker stop cart-service
	@docker rm -f cart-service
	@docker rmi -f cart-service:0.0.1