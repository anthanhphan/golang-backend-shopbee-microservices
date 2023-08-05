run:
	@node server.js

deploy:
	@docker rmi -f api-gateway:0.0.1
	@docker compose up -d

rm:
	@docker stop api-gateway
	@docker rm -f api-gateway
	@docker rmi -f api-gateway:0.0.1