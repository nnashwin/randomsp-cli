all: build move
	

build:
	go build -o rsp

move:
	mv rsp $(GOPATH)/bin
