.PHONY: run air

run:
	export $$(cat .env |xargs) && go run cmd/server/main.go

air: 
	go build -o tmp/main cmd/server/main.go

air_run:
	docker-compose -f air.yaml down && \
	docker-compose -f air.yaml up --build -d --remove-orphans && \
	docker logs air -f