BINARY_NAME := go-tawerin
IMG ?= go-tawerin
ROOT_PACKAGE := github.com/cagiti/go-tawerin

os = $(word 1, $@)

BUILD_DIR= $(CURDIR)/build

REV        := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
VERSION    ?= $(shell echo "$$(git describe --abbrev=0 --tags 2>/dev/null)-dev-$(REV)")
COMMIT     := $(shell git rev-parse HEAD 2> /dev/null || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)

BUILDFLAGS := -ldflags \
   "-s -w \
    -X $(ROOT_PACKAGE)/internal/version.Version=$(VERSION)\
	-X $(ROOT_PACKAGE)/internal/version.Commit=$(COMMIT)\
	-X $(ROOT_PACKAGE)/internal/version.BuildDate=$(BUILD_DATE)"

$(BUILD_DIR): ## Creates the output directory for all build related outputs
	@mkdir -p $(BUILD_DIR)

.PHONY: $(all)
all: linux test check ## Builds the Linux binary, runs tests and linters

.PHONY: linux
linux: $(BUILD_DIR) ## Build the OS specific binary
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./web/tawerin/tawerin.go

.PHONY: test
test: $(BUILD_DIR) ## Runs the tests
	go test ./... -coverprofile $(BUILD_DIR)/cover.out

check: fmt vet ## Runs all linters

fmt: ## Runs go fmt
	go fmt ./...

vet: ## Runs go vet
	go vet ./...

tidy: ## Cleans up these unused dependencies
	go mod tidy

clean: ## Deletes the build output directory
	rm -rf $(BUILD_DIR)

docker-build: linux ## Build the docker image
	docker build --network host -t $(IMG):$(VERSION) .

.PHONY: help
help: ## Prints this help
