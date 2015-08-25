#!/bin/bash

# help #
help:
	@echo "\033[32mAvailable commands\033[0m"
	@echo "\033[33mAPI\033[0m"
	@echo "- \033[1mapi-run:\033[0m Run the application"
	@echo "- \033[1mapi-run-command:\033[0m Run a command (eg: make api-run-command COMMAND=\"run-app.sh\")"
	@echo "- \033[1mapi-run-go-command:\033[0m Run a command (eg: make api-run-command COMMAND=\"hello_world.go\")"
	@echo "- \033[1mapi-fmt:\033[0m Format the code"
	@echo "- \033[1mapi-fmt-check:\033[0m Check code format"
	@echo "- \033[1mapi-test:\033[0m Run application tests"
	@echo "\033[33mDocker compose\033[0m"
	@echo "- \033[1mcmp-up:\033[0m Up containers"
	@echo "- \033[1mcmp-pull:\033[0m Update containers"
	@echo "- \033[1mcmp-stop:\033[0m Stop containers"
	@echo "\033[33mMongoDB\033[0m"
	@echo "- \033[1mmongo-connect:\033[0m Open mongodb term (container must be already run)"
	@echo ""

# api #
api-run:
	docker-compose run --rm --service-ports api ./bin/run.sh ./bin/run-app.sh

api-run-command:
	docker-compose run --rm api ./bin/run.sh ./bin/$(COMMAND)

api-run-go-command:
	docker-compose run --rm api ./bin/run.sh "go run ./bin/$(COMMAND).go"

api-fmt:
	docker run --rm -v `pwd`:/go golang:1.4 gofmt -w /go

api-fmt-check:
	docker run --rm -v `pwd`:/go golang:1.4 gofmt -d /go

api-test:
	echo "Not Implemented"

# docker-compose # 
cmp-up:
	docker-compose up -d

cmp-pull:
	docker-compose pull

cmp-stop:
	docker-compose stop

# mongodb #
mongo-connect:
	docker exec -t -i api_database_1  /bin/bash
