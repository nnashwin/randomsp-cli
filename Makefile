all: build move

build:
	go build -o rsp

move:
	if [ "$(GOPATH)" = "" ] ; \
	then \
		echo "Your GOPATH is not set.  You must set it to continue"; \
	else \
		mv rsp $(GOPATH)/bin/; \
	fi;
