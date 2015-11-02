export PATH := ../..:${PATH}

build: flag.deps
	go build -o example

test: message generate build
	go test -v

clean:
	rm -f *_store.go
	rm -f *_endpoints.go
	rm -f *_rest.go
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
	go get github.com/gourd/codec
	go get github.com/gourd/kit/store
	go get github.com/gourd/kit/store/upperio
	go get github.com/gourd/kit/oauth2
	go get github.com/gourd/kit/perm
	go get github.com/yookoala/restit
	touch flag.deps

.PHONY: build test clean generate message
