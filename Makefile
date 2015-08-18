export PWD=$(shell pwd)
export PATH:=${PWD}:${PATH}

build: flag.deps
	go build -o gourd

test: build flag.testDeps
	@echo
	@echo "============================="
	@echo " Run Examples as Test"
	@echo "============================="
	cd _examples/1_basic && make test
	cd _examples/2_server && make test
	cd _examples/3_string_id && make test
	cd _examples/4_service_unavailable && make test

clean:
	rm -f gourd
	rm -f flag.*
	cd _examples/1_basic && make clean
	cd _examples/2_server && make clean
	cd _examples/3_string_id && make clean
	cd _examples/4_service_unavailable && make clean


flag.deps:
	go get -u github.com/gourd/service/upperio
	go get -u upper.io/db/sqlite
	go get -u github.com/codegangsta/cli
	go get -u github.com/gourd/goparser
	touch flag.deps

flag.testDeps:
	go get -u github.com/codegangsta/negroni
	go get -u github.com/gorilla/pat
	go get -u github.com/satori/go.uuid
	go get -u github.com/gourd/codec
	go get -u github.com/gourd/oauth2
	go get -u github.com/gourd/perm
	go get -u github.com/yookoala/restit
	touch flag.testDeps

.PHONY: build test clean