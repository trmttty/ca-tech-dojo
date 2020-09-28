DOCKER_COMPOSE_FILE='deployment/docker-compose.yml'

all: up

build:
	docker-compose -f ${DOCKER_COMPOSE_FILE} build

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d

up-b:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d --build

down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down

.PHONY: build up down