export PATH := ../..:${PATH}

build: prebuild
	echo "gourd path: `which gourd`"
	go generate

prebuild: flag.get_deps

flag.get_deps:
	go get -u github.com/codegangsta/negroni
	go get -u github.com/gourd/service
	go get -u github.com/gourd/service/upperio
	go get -u github.com/gourd/codec
	go get -u github.com/gourd/oauth2
	go get -u github.com/gourd/perm
	go get -u github.com/gorilla/pat
	go get -u upper.io/db/postgresql
	go get -u github.com/yookoala/restit
	touch flag.get_deps

test: build
	go test -v

clean:
	rm -f *_service.go
	rm -f *_service_rest.go
	rm -f flag.*

.PHONY: build prebuild test clean
