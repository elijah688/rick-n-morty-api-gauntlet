.PHONY: db_up run migration server gateway run ui

db_up:
	docker-compose up -d

migration: db_up
	export $$(cat .env |xargs) && cd ./go/server && go run cmd/migration/main.go


server:
	cd ./go/server && make air_run


gateway:
	cd ./go/gateway && make air_run


ui:
	cd ./ui/my-react-app; npm run dev

run: db_up
	(make migration) && \
	(make server &) && \
	(make gateway &) && \
	(make ui &)

