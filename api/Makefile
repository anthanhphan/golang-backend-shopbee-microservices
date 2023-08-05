run:
	@cd cmd;go run main.go

deploy:
	@docker rmi -f shopbee-api:0.0.1
	@docker compose up -d

rm:
	@docker stop shopbee-api
	@docker rm -f shopbee-api
	@docker rmi -f shopbee-api:0.0.1