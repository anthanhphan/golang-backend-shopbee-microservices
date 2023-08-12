run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f admin-service:0.0.1
	@docker compose up -d

rm:
	@docker stop admin-service
	@docker rm -f admin-service
	@docker rmi -f admin-service:0.0.1