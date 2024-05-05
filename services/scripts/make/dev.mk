.PHONY: status logs start stop clean start-service

status: ## Show status of docker containers
	docker-compose ps

logs: ## Show logs of docker containers
	docker-compose logs -f


stop: ## Stop docker containers
	docker-compose stop

# clean:stop ## Stop docker containers, clean data and workspace
# 	docker-compose down -v --remove-orphans
# 	docker rmi services-user

start-service:
	docker-compose -f ./$(SERVICE)/docker-compose.yml up --build

stop-service:
	docker-compose -f ./$(SERVICE)/docker-compose.yml down
