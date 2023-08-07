run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f upload-service:0.0.1
	@docker compose up -d

rm:
	@docker stop upload-service
	@docker rm -f upload-service
	@docker rmi -f upload-service:0.0.1