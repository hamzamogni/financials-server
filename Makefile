.DEFAULT_GOAL := build

migrate:
	docker-compose run app go run app/main.go migrate

run:
	docker-compose up --build
.PHONY:run
