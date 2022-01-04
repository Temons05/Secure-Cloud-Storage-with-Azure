SHELL := /bin/bash

build:
	cd frontend && yarn install && yarn build
	Copy-Item -path frontend/dist/ -Destination static
	go mod download
	pkger
	go build -o webserver

docker-build:
	docker build -t azurewebserver.azurecr.io/webserver .

docker-run: docker-build
	docker run -p 8081:8081 azurewebserver.azurecr.io/webserver

docker-push: 
	docker push azurewebserver.azurecr.io/webserver

run: build
	chmod +x ./webserver
	./webserver