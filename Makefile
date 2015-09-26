.PHONY: all \
        clean \
        cleanall \
        dev \
        help \
        lint \
        release \
        test \
        testall

PROJECT = packer-post-processor-vhd
VERSION = 0.3.0

all: $(PROJECT)

$(PROJECT): clean
	go build .

help:
	@echo "clean     remove testing artifacts"
	@echo "cleanall  remove development and testing artifacts"
	@echo "dev       set up development environment"
	@echo "help      show this page"
	@echo "lint      check style with golint"
	@echo "test      run unit tests"
	@echo "testall   run integration tests"
	@echo "release   push tags upstream"

clean:
	go clean -x
	$(RM) -r test/output-virtualbox-iso-o*
	$(RM) test/*.vhd

cleanall: clean
	$(RM) -r test/output-virtualbox-iso

dev:
	cd test && packer build --force fixtures/virtualbox-iso.json

lint:
	golint ./...

release:
	git push origin --tags

test:
	go test -v ./...

testall:
	go test -tags=integration -v ./...
