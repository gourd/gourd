export PATH := ../..:${PATH}

build: flag.deps
	go build

test: message generate build
	go test -v

clean:
	rm -f *_service.go
	rm -f *_service_rest.go
	rm -f flag.*


generate:
	@echo "generateing with gourd: `which gourd`"
	go generate

message:
	@echo
	@echo "-----------------------------"
	@echo " Example 1"
	@echo " - Simple file generation"
	@echo "-----------------------------"


flag.deps:
	go get github.com/gourd/service
	touch flag.deps

.PHONY: build test clean generate message
