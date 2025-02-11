run:
	go run main.go

build:
	docker build -t chat-app .

up:
	docker-compose up --build

migrate:
	psql -h localhost -U postgres -d chatapp -f database/migrate.sql
