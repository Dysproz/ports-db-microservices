help:               ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:              ## builds clientAPI and portDomainService containers with docker-compose
	@docker-compose build --no-cache

up:                 ## creates all containers with docker-compose
	@docker-compose up --force-recreate

load-ports:         ## loads ports file using clientAPI HTTP endpoint. Default file is ports.json, custom one may be set with file variable - Example usage: make load-ports file=myports.json
	@if [ "$(file)" = "" ]; then \
		curl -F file=@ports.json 'http://127.0.0.1:8000/loadPorts'; \
	else \
		curl -F file=@$(file) 'http://127.0.0.1:8000/loadPorts'; \
	fi;

get:                ## gets a port data from database using clientAPI HTTP endpoint. Example usage: make get key=ZWUTA
	@if [ "$(key)" = "" ]; then \
		echo "No key variable passed. Example usage: make get key=ZWUTA"; \
	else \
		curl -X POST -H "Content-Type: application/json" -d '{"key": "$(key)"}'  'http://127.0.0.1:8000/getPort'; \
	fi;
test:               ## executes Go unit tests
	@go test ./internal/...

lint:               ## Lints Go code
	@golint ./...

fmt:                ## Format Go code
	@go fmt ./...

protogen:           ## Generates Go code based on .proto files
	@protoc --go_out=. --go_opt=module=github.com/Dysproz/ports-db-microservices --go-grpc_out=. --go-grpc_opt=module=github.com/Dysproz/ports-db-microservices *.proto
