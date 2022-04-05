build:
	docker-compose build --no-cache

up:
	docker-compose up --force-recreate

load-ports:
	curl -F file=@ports.json 'http://127.0.0.1:8000/loadPorts'

