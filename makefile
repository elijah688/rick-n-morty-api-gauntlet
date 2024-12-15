.PHONY: db_up run migration server gateway

db_up:
	docker-compose up -d

migration: db_up
	export $$(cat .env |xargs) && cd ./go/server && go run cmd/migration/main.go


server:
	cd ./go/server && make air_run


gateway:
	cd ./go/gateway && make air_run
