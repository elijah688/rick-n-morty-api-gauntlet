.PHONY: db_up run migration server gateway run ui kill go_server go_gateway go_run py_server py_gateway py_run 


kill:
	lsof -i:8080 -t | xargs kill -9
	lsof -i:8081 -t | xargs kill -9
	lsof -i:5174 -t | xargs kill -9

db_up:
	docker-compose up -d

migration: db_up
	export $$(cat .env |xargs) && cd ./go/server && go run cmd/migration/main.go


ui:
	cd ./ui/my-react-app; npm run dev

go_server:
	cd ./go/server && make air_run


go_gateway:
	cd ./go/gateway && make air_run


go_run: kill db_up
	(make migration) && \
	(make go_server &) && \
	(make go_gateway &) && \
	(make ui &)


py_server:
	cd python/server && \
	. ./venv/bin/activate && \
	pip install --upgrade pip && \
	pip install --upgrade -r requirements.txt && \
	cd app && \
	export $$(cat .env | xargs) && \
	python3 main.py

py_gateway:
	cd python/gateway && \
	. ./venv/bin/activate && \
	pip install --upgrade pip && \
	pip install --upgrade -r requirements.txt && \
	cd app && \
	export $$(cat .env | xargs) && \
	python3 main.py


py_run: db_up kill
	(make migration) && \
	(make py_server &) && \
	(make py_gateway &) && \
	(make ui &)