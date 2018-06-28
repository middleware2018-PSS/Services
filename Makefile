# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help build-up up down

help: ## Shows this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


# DOCKER COMPOSE TASKS
gen-keys: ## Generate the RSA key pair needed by JWT
	openssl genrsa -out config/back2school.rsa 1024 && openssl rsa -in config/back2school.rsa -pubout > config/back2school.rsa.pub

build-up: ## Build images before starting containers with up
	sudo docker-compose up --build

build-up-detached: ## Build images before starting containers with up (detached mode)
	sudo docker-compose up --build -d

up: ## Builds, (re)creates, starts, and attaches to containers for a service
	sudo docker-compose up

up-detached: ## Builds, (re)creates, starts, and attaches to containers for a service (detached mode)
	sudo docker-compose up -d

down: ## Stops containers and removes containers, networks, volumes, and images created by up
	sudo docker-compose down

db-init: ## Restore the dump of the database with the initial data
	sudo docker-compose exec db psql -U postgres -f /dump.sql
