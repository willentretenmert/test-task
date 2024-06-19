.PHONY: up down bench

all: build up-db fill up-app bench

build:
	@echo "building containers..."
	docker-compose build

up-db:
	@echo "starting db container..."
	docker-compose up -d db

up-app:
	@echo "starting app container..."
	docker-compose up -d app

down:
	@echo "stopping containers..."
	docker-compose down

fill:
	@echo "filling up database..."
	cd tools && go mod tidy && go run csv-exporter.go

bench:
	@echo "running k6 test..."
	k6 run tools/script.js