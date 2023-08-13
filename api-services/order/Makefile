run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f order-service:0.0.1
	@docker compose up -d

rm:
	@docker stop order-service
	@docker rm -f order-service
	@docker rmi -f order-service:0.0.1