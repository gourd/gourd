export PATH := ../..:${PATH}

build: prebuild
	echo "gourd path: `which gourd`"
	go generate

prebuild: flag.get_deps
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql

flag.get_deps:
	go get -u github.com/gourd/service
	go get -u github.com/gourd/service/upperio
	go get -u upper.io/db/sqlite
	touch flag.get_deps

test: build
	go test -v

clean:
	rm -f *_service.go
	rm -f flag.*

.PHONY: build prebuild test clean
