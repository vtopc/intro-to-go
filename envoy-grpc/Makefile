.PHONY: up
up:
	docker-compose up

.PHONY: re-build
re-build: build-grpc-server
	docker-compose up --build

.PHONY: build-grpc-server
build-grpc-server:
	cd ./grpc-server && $(MAKE) protos deps
