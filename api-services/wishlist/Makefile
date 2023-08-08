run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f wishlist-service:0.0.1
	@docker compose up -d

rm:
	@docker stop wishlist-service
	@docker rm -f wishlist-service
	@docker rmi -f wishlist-service:0.0.1