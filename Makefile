# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = test-with-coverage

SHELL = /bin/bash -euo pipefail

GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   = go.octolab.net
GOPROXY     = direct
LOCAL       = $(MODULE)
MODULE      = `go list -m`
PACKAGES    = `go list ./... 2> /dev/null`
PATHS       = $(shell echo $(PACKAGES) | sed -e "s|$(MODULE)/\{0,1\}||g")
TIMEOUT     = 1s

ifeq (, $(PACKAGES))
	PACKAGES = $(MODULE)
endif

ifeq (, $(PATHS))
	PATHS = .
endif

export GO111MODULE := $(GO111MODULE)
export GOFLAGS     := $(GOFLAGS)
export GOPRIVATE   := $(GOPRIVATE)
export GOPROXY     := $(GOPROXY)

.PHONY: go-env
go-env:
	@echo "GO111MODULE: `go env GO111MODULE`"
	@echo "GOFLAGS:     $(strip `go env GOFLAGS`)"
	@echo "GOPRIVATE:   $(strip `go env GOPRIVATE`)"
	@echo "GOPROXY:     $(strip `go env GOPROXY`)"
	@echo "LOCAL:       $(LOCAL)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"

.PHONY: deps
deps:
	@go mod tidy
	@if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi

.PHONY: deps-clean
deps-clean:
	@go clean -modcache

.PHONY: update
update: selector = '.Require[] | select(.Indirect != true) | .Path'
update:
	@if command -v egg > /dev/null; then \
		packages="`egg deps list`"; \
		go get -mod= -u $$packages; \
	elif command -v jq > /dev/null; then \
		packages="`go mod edit -json | jq -r $(selector)`"; \
		go get -mod= -u $$packages; \
	else \
		packages="$(shell cat go.mod | grep -v '// indirect' | grep '\t' | awk '{print $$1}')"; \
		go get -mod= -u $$packages; \
	fi

BINARY  = $(BINPATH)/$(shell basename $(MAIN))
BINPATH = $(PWD)/bin
COMMIT  = $(shell git rev-parse --verify HEAD)
DATE    = $(shell date +%Y-%m-%dT%T%Z)
LDFLAGS = -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)"
MAIN    = $(MODULE)

export GOBIN := $(BINPATH)
export PATH  := $(BINPATH):$(PATH)

.PHONY: build-env
build-env:
	@echo "BINARY:      $(BINARY)"
	@echo "BINPATH:     $(BINPATH)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "DATE:        $(DATE)"
	@echo "GOBIN:       `go env GOBIN`"
	@echo "LDFLAGS:     $(LDFLAGS)"
	@echo "MAIN:        $(MAIN)"
	@echo "PATH:        $$PATH"

.PHONY: build
build:
	@go build -o $(BINARY) $(LDFLAGS) $(MAIN)

.PHONY: build-clean
build-clean:
	@go clean -cache

.PHONY: install
install: build

.PHONY: install-clean
install-clean:
	@rm -f $(BINARY)

.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-clean
test-clean:
	@go clean -testcache

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > bin/install

.PHONY: format
format:
	@goimports -local $(LOCAL) -ungroup -w $(PATHS)

.PHONY: generate
generate:
	@go generate $(PACKAGES)


.PHONY: clean
clean: build-clean deps-clean install-clean test-clean

.PHONY: env
env: go-env build-env

.PHONY: refresh
refresh: update deps generate format test build
