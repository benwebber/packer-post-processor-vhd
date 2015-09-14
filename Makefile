.PHONY: all \
        clean \
        help \
        lint \
        release \
        test \
        testall

PROJECT = packer-post-processor-vhd
VERSION = 0.2.0

all: $(PROJECT)

$(PROJECT): clean
	go build .

help:
	@echo "clean    remove build artifacts"
	@echo "help     show this page"
	@echo "lint     check style with golint"
	@echo "test     run unit tests"
	@echo "testall  run integration tests"
	@echo "release  push tags upstream"

clean:
	go clean -x

lint:
	golint ./...

release:
	git push origin --tags

test:
	go test -v ./...

testall:
	go test -tags=integration -v ./...
