all: install

GOPATH:=$(CURDIR)
export GOPATH

fmt:
	gofmt -l -w -s src/

dep:fmt
	go install checker
	go install git
	go install command
	go install lab

install:dep
