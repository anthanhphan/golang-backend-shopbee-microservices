run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f product-service:0.0.1
	@docker compose up -d

rm:
	@docker stop product-service
	@docker rm -f product-service
	@docker rmi -f product-service:0.0.1