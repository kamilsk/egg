# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = install

SHELL = /bin/bash -euo pipefail

GO    = GOPRIVATE=$(GOPRIVATE) GOFLAGS=$(GOFLAGS) go

BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   = go.octolab.net
GOPROXY     = direct
MODULE      = $(shell $(GO) list -m)

export PATH := $(BINPATH):$(PATH)

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(shell $(GO) env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell $(GO) env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell $(GO) env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell $(GO) env GOPROXY))"
	@echo "MODULE:      $(MODULE)"

.PHONY: deps
deps:
	@$(GO) mod tidy
	@if [[ "$(shell $(GO) env GOFLAGS)" =~ -mod=vendor ]]; then $(GO) mod vendor; fi

.PHONY: deps-clean
deps-clean:
	@$(GO) clean -modcache

.PHONY: update
update: selector = '.Require[] | select(.Indirect != true) | .Path'
update:
	@if command -v egg > /dev/null; then \
		packages="$(shell egg deps list)"; \
		$(GO) get -mod= -u $$packages; \
	elif command -v jq > /dev/null; then \
		packages="$(shell $(GO) mod edit -json | jq -r $(selector))"; \
		$(GO) get -mod= -u $$packages; \
	else \
		packages="$(shell cat go.mod | grep -v '// indirect' | grep '\t' | awk '{print $$1}')"; \
		$(GO) get -mod= -u $$packages; \
	fi

.PHONY: build
build: install

.PHONY: build-clean
build-clean:
	@$(GO) clean -cache

.PHONY: install
install:
	@ROOT=$(dir $(BINPATH)) $(GO) generate tools.go

.PHONY: install-clean
install-clean:
	@if command -v egg > /dev/null; then \
		tools="$(shell egg tools list)"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	else \
		tools="$(shell cat tools.go | grep 'go:generate' | awk '{print $$NF}' | xargs basename)"; \
		for tool in $$tools; do rm -f $(BINPATH)/$$tool; done; \
	fi


.PHONY: clean
clean: build-clean deps-clean install-clean
