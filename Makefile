.DEFAULT_GOAL := build

migrate:
	docker-compose run app go run app/main.go migrate

seed:
	go run app/main.go seed

run:
	docker-compose up --build
.PHONY:run
