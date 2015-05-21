# api #
api-run:
	docker-compose run api

api-fmt:
	@docker run --rm -v `pwd`/src:/go/src golang:1.4 gofmt -w /go/src

api-fmt-check:
	@docker run --rm -v `pwd`/src:/go/src golang:1.4 gofmt -d /go/src

api-test:
	echo "Not Implemented!!!"

# docker-compose # 
cmp-up:
	docker-compose up -d

cmp-pull:
	docker-compose pull

cmp-stop:
	docker-compose stop