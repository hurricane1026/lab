all: install

GOPATH:=$(CURDIR)
export GOPATH

fmt:
	gofmt -l -w -s src/

dep:fmt
	go get github.com/hurricane1026/gcfg

install:dep
	go install checker
	go install git
	go install command
	go install lab
