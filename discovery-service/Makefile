run:
	@node server.js

deploy:
	@docker rmi -f discovery-service:0.0.1
	@docker compose up -d

rm:
	@docker stop discovery-service
	@docker rm -f discovery-service
	@docker rmi -f discovery-service:0.0.1