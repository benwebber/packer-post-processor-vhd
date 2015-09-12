.PHONY: all

PROJECT = packer-post-processor-vhd
VERSION = 0.1.0

all: build

clean:
	$(RM) $(PROJECT)

build:
	go build .

release:
	git push origin --tags
