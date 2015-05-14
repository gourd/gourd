build: prebuild
	go generate

prebuild:
	go get github.com/codegangsta/negroni
	go get github.com/gourd/service
	go get github.com/gorilla/pat
	go get upper.io/db/sqlite
	go get github.com/yookoala/restit
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql

test: build
	go test

.PHONY: build prebuild test
