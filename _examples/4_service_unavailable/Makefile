export PATH := ../..:${PATH}

build: flag.deps
	go build -o example

test: message generate build
	go test -v

clean:
	rm -f *_service.go
	rm -f *_service_rest.go
	rm -f flag.*
	rm -f example


generate:
	@echo "gourd path: `which gourd`"
	go generate

message:
	@echo
	@echo "-----------------------------"
	@echo " Example 4"
	@echo " - A dummy API server build"
	@echo "-----------------------------"


flag.deps:
	go get github.com/codegangsta/negroni
	go get github.com/gorilla/pat
	go get upper.io/db/postgresql
	go get github.com/gourd/service
	go get github.com/gourd/service/upperio
	go get github.com/gourd/codec
	go get github.com/gourd/oauth2
	go get github.com/gourd/perm
	go get github.com/yookoala/restit
	touch flag.deps

.PHONY: build test clean generate message
