.PHONY: all

PROJECT = packer-post-processor-vhd
VERSION = 0.1.0

all: build

build: clean
	go build .

clean:
	go clean -x

release:
	git push origin --tags

test:
	go test -v ./...
