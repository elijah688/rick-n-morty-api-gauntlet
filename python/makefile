.PHONY: run server gateway

server:
	cd server && \
	. ./venv/bin/activate && \
	pip install -r requirements.txt \
	cd app && \
	export $$(cat .env | xargs) && \
	python3 main.py


gateway:
	cd gateway && \
	. ./venv/bin/activate && \
	pip install -r requirements.txt \
	cd app && \
	export $$(cat .env | xargs) && \
	python3 main.py


run:
	make server & \
	make gateway & \