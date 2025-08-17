-include .env

# Redefinir MAKEFILE_LIST pour qu'il ne contienne que le Makefile
MAKEFILE_LIST := Makefile

ENV_FILE := --env-file .env

# Couleurs
GREEN = \033[0;32m
YELLOW = \033[0;33m
NC = \033[0m # No Color

# Variables
COMPOSE_FILE = $(if $(filter $(APP_ENV),prod),docker/compose.prod.yaml,$(if $(filter $(APP_ENV),preprod),docker/compose.preprod.yaml,docker/compose.yaml))
DOCKER_COMPOSE = docker compose $(ENV_FILE) -f $(COMPOSE_FILE)

.PHONY: help build up down logs shell restart clean status ps ta tap tav tavp tf tfv

help: ## Affiche cette aide
	@echo ""
	@echo "Commandes disponibles:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) $(YELLOW)%s$(NC)\n", $$1, $$2}'
	@echo ""

build: ## Construit toutes les images Docker
	$(DOCKER_COMPOSE) build

build-app: ## Construit uniquement l'image de l'application frontend
	$(DOCKER_COMPOSE) build app

build-api: ## Construit uniquement l'image de l'API backend
	$(DOCKER_COMPOSE) build api

up: ## Lance tous les services
	$(DOCKER_COMPOSE) up -d

upb: ## rebuild l'image et lance tous les services
	$(DOCKER_COMPOSE) up -d --build

up-logs: ## Lance tous les services avec les logs
	$(DOCKER_COMPOSE) up

down: ## Arrête tous les services et supprime les containers
	$(DOCKER_COMPOSE) down

stop: ## Arrête tous les services
	$(DOCKER_COMPOSE) stop

cm: ## créé un fichier de migration - usage: make cm file=nom_du_fichier
	docker exec -it tenjin_$(APP_ENV)_api migrationcreate $(file)

logs: ## Affiche les logs de tous les services
	$(DOCKER_COMPOSE) logs -f

logs-app: ## Affiche les logs de l'application frontend
	$(DOCKER_COMPOSE) logs -f app

logs-api: ## Affiche les logs de l'API backend
	$(DOCKER_COMPOSE) logs -f api

logs-db: ## Affiche les logs de la base de données
	$(DOCKER_COMPOSE) logs -f db

shell-app: ## Ouvre un shell dans le conteneur de l'application frontend
	$(DOCKER_COMPOSE) exec app sh

shell-api: ## Ouvre un shell dans le conteneur de l'API backend
	$(DOCKER_COMPOSE) exec api bash

shell-db: ## Ouvre un shell dans le conteneur de la base de données
	$(DOCKER_COMPOSE) exec db mongosh

clean: ## Supprime les conteneurs, réseaux et volumes
	$(DOCKER_COMPOSE) down -v --remove-orphans

clean-all: ## Supprime tout (conteneurs, réseaux, volumes et images)
	$(DOCKER_COMPOSE) down -v --remove-orphans --rmi all

ta: ## Lance tous les tests api
	docker exec -i -e APP_ENV=test tenjin_dev_api go test ./...

tai: ## Lance tous les tests api d'integration avec logs (fmt-print)
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -tags=integration ./...

tap: ## Lance les tests api pour un path spécifique (usage: make tap path=monpath)
	docker exec -i -e APP_ENV=test tenjin_dev_api go test ./$(path)

taip: ## Lance les tests api d'integration pour un path spécifique (usage: make tap path=monpath)
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -tags=integration ./$(path)

tav: ## Lance tous les tests api en verbose
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -v ./...

taiv: ## Lance tous les tests api en verbose + integration
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -tags=integration -v ./...

tavp: ## Lance les tests api en verbose pour un path (usage: make tavp path=monpath)
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -v ./$(path)

taivp: ## Lance les tests api en verbose + integration pour un path (usage: make tavp path=monpath)
	docker exec -i -e APP_ENV=test tenjin_dev_api go test -v -tags=integration ./$(path)

tf: ## Lance tous les tests front
	docker exec -i tenjin_dev_app npm run test.unit -- --run

tfv: ## Lance tous les tests front en mode verbose
	docker exec -i tenjin_dev_app npm run test.unit -- --run --reporter=verbose

cli: ## execute commande package api rome (usage: make apir cmd=help)
	$(DOCKER_COMPOSE) exec api bash -c "go run ./cli/main.go $(cmd)"
