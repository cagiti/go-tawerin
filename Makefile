all: release

clean:
	rm -rf go-tawerin Godeps vendor

install: clean prepare build
	godep go install

prepare: clean
	go get github.com/tools/godep
	go get github.com/gorilla/mux
	go get github.com/newrelic/go-agent

build: clean prepare
	godep save
	godep go build

test: clean prepare build install
	echo "no unit tests"
	go vet .

release: clean prepare build install test

.PHONY: clean install prepare build test release
