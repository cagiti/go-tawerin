all:
  test

clean:
  rm -f go-tawerin

install:
  prepare
  godep go install

prepare:
  go get github.com/tools/godep
  go get github.com/gorilla/mux
  go get github.com/newrelic/go-agent

build:
  prepare
  godep go build

test:
  prepare
  build
  echo "no tests"

.PHONY: install prepare build test
