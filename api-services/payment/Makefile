run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f payment-service:0.0.1
	@docker compose up -d

rm:
	@docker stop payment-service
	@docker rm -f payment-service
	@docker rmi -f payment-service:0.0.1