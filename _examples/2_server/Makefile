export PATH := ../..:${PATH}

build: generate
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql
	go build

generate:
	echo "generateing with gourd: `which gourd`"
	go generate

flag.get_deps:
	go get -u github.com/codegangsta/negroni
	go get -u github.com/gourd/service
	go get -u github.com/gourd/service/upperio
	go get -u github.com/gourd/codec
	go get -u github.com/gourd/oauth2
	go get -u github.com/gourd/perm
	go get -u github.com/gorilla/pat
	go get -u upper.io/db/sqlite
	go get -u github.com/yookoala/restit
	touch flag.get_deps

test: flag.get_deps generate
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql
	go test -v

clean:
	rm -f *_service.go
	rm -f *_service_rest.go
	rm -f flag.*

.PHONY: build generate test clean
