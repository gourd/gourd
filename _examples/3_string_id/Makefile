export PATH := ../..:${PATH}

build: flag.deps
	go build

test: message generate build
	sqlite3 data/sqlite3.db < data/schema.sqlite3.sql
	go test -v

clean:
	rm -f *_service.go
	rm -f flag.*


generate:
	@echo "gourd path: `which gourd`"
	go generate

message:
	@echo
	@echo "------------------------------------------"
	@echo " Example 3"
	@echo " - Simple file generation with string ID"
	@echo "------------------------------------------"


flag.deps:
	go get github.com/gourd/kit/store
	go get github.com/gourd/kit/store/upperio
	go get upper.io/db/sqlite
	touch flag.deps

.PHONY: build test clean generate message
