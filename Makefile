include .env

SHELL=/bin/bash
APP=app
export SERVICE=sentry_log
export COMPOSE_PROJECT_NAME=20
.DEFAULT_GOAL := up


init: down
	@go mod init mod && go mod tidy

build:
	@go build -o ${APP} *.go 

up: build
	@./${APP}

down:
	@[ -f go.mod ] && bash -c 'rm go.mod go.sum' || echo go.mod go.sum are not exist
	@[ -f ${APP} ] && rm ${APP} || echo ${APP} are not exist

.PHONY: all

all: init build up

docker-build:
	docker build --no-cache -t ${SERVICE} .

docker-run: docker-down
	@for i in $(shell seq 1 ${COMPOSE_PROJECT_NAME});do COMPOSE_PROJECT_NAME=$$i docker compose up -d ;done

docker-down:
	@for i in $(shell seq 1 ${COMPOSE_PROJECT_NAME});do COMPOSE_PROJECT_NAME=$$i docker compose down;done

d-all: docker-build docker-down docker-run