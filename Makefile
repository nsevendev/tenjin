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

.PHONY: help build up down logs shell restart clean status ps

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
