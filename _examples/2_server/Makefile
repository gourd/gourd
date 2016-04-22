export PATH := ../..:${PATH}

build: flag.deps
	go build -o example

test: message generate build
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql
	go test -v

clean:
	rm -f *_store.go
	rm -f *_endpoints.go
	rm -f *_rest.go
	rm -f flag.*
	rm -f example


generate:
	@echo "generateing with gourd: `which gourd`"
	go generate

message:
	@echo
	@echo "-----------------------------"
	@echo " Example 2"
	@echo " - A dummy API server build"
	@echo "-----------------------------"


flag.deps:
	go get github.com/codegangsta/negroni
	go get github.com/gorilla/pat
	go get upper.io/db/sqlite
	go get github.com/gourd/kit/store
	go get github.com/gourd/kit/store/upperio
	go get github.com/gourd/kit/oauth2
	go get github.com/gourd/kit/perm
	go get github.com/yookoala/restit/v1
	touch flag.deps

.PHONY: build test clean generate message
