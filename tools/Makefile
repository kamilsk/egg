# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = install

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

BINPATH = $(shell dirname $(PWD))/bin

export GOBIN := $(BINPATH)
export PATH  := $(BINPATH):$(PATH)

.PHONY: tools-env
tools-env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GOBIN:       `go env GOBIN`"
	@echo "PATH:        $$PATH"

.PHONY: build
build: install

.PHONY: build-clean
build-clean:
	@go clean -cache

.PHONY: install
install:
	@go generate tools.go

.PHONY: install-clean
install-clean:
	@if command -v egg > /dev/null; then \
		tools="`egg tools list`"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	else \
		tools="$(shell cat tools.go | grep 'go:generate' | awk '{print $$NF}' | xargs basename)"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	fi


.PHONY: clean
clean: build-clean deps-clean install-clean

.PHONY: env
env: go-env tools-env

.PHONY: refresh
refresh: update deps install
