.PHONY: googleapis
googleapis:
	rm -rf /tmp/googleapis
	git clone https://github.com/googleapis/googleapis.git /tmp/googleapis
	cp -R /tmp/googleapis/google protorepo/

.PHONY: protos
protos:
	docker run -v `pwd`:/workspace --rm grpckit/omniproto:1.32_0
	#docker run -v `pwd`:/defs --rm namely/protoc-all -d protos/helloworld -l go

.PHONY: deps
deps:
	go mod download -json

.PHONY: test
test:
	go test `go list ./... | grep -v '/gen' | grep -v '/protos'` -cover -count=1 -coverprofile=coverage.txt -covermode=count

.PHONY: run
run:
	go run ./cmd/hello
