
# help #
help:
	@echo "Available commands: "
	@echo "API:"
	@echo "- api-run: Run the application"
	@echo "- api-fmt: Format the code"
	@echo "- api-fmt-check: Check code format"
	@echo "- api-test: Run application tests"
	@echo ""
	@echo "Docker compose:"
	@echo "- cmp-up: Up containers"
	@echo "- cmp-pull: Update containers"
	@echo "- cmp-stop: Stop containers"
	@echo ""
	@echo "MongoDB"
	@echo "- mongo-connect: Open mongodb term (container must be already run)"
	@echo ""
# api #
api-run:
	docker-compose run api

api-fmt:
	docker run --rm -v `pwd`/src:/go/src golang:1.4 gofmt -w /go/src

api-fmt-check:
	docker run --rm -v `pwd`/src:/go/src golang:1.4 gofmt -d /go/src

api-test:
	echo "Not Implemented"

# docker-compose # 
cmp-up:
	docker-compose up -d

cmp-pull:
	docker-compose pull

cmp-stop:
	docker-compose stop

#mongodb#
mongo-connect:
	docker exec -t -i api_database_1  /bin/bash