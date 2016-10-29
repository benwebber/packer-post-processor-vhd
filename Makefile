.PHONY: .depend \
        all \
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

$(PROJECT): .depend clean
	go build -v .

help:
	@echo "clean     remove testing artifacts"
	@echo "cleanall  remove development and testing artifacts"
	@echo "dev       set up development environment"
	@echo "dist      cross-compile binaries for distribution"
	@echo "help      show this page"
	@echo "lint      check style with golint"
	@echo "test      run unit tests"
	@echo "testall   run integration tests"
	@echo "release   push tags and binaries upstream"

.depend:
	go get -d github.com/mitchellh/packer

clean:
	go clean -x
	$(RM) -r test/output-virtualbox-iso-o*
	$(RM) test/*.vhd

cleanall: clean
	$(RM) -r test/output-virtualbox-iso

dev:
	cd test && packer build --force fixtures/virtualbox-iso.json

dist:
	gox --osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/$(PROJECT)-$(VERSION)-{{ .OS }}_{{ .Arch }}"

lint:
	golint ./...

release: dist
	git push origin --tags
	scripts/release.sh $(PROJECT) $(VERSION)

test:
	go test -v ./...

testall:
	install $(PROJECT) test/
	go test -tags=integration -v ./...
