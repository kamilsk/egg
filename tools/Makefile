.DEFAULT_GOAL = install

SHELL = /bin/bash -euo pipefail

BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   =
GOPROXY     = direct
MODULE      = $(shell go list -m)

export PATH := $(BINPATH):$(PATH)

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell go env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell go env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell go env GOPROXY))"
	@echo "MODULE:      $(MODULE)"

.PHONY: deps
deps:
	@GOPRIVATE=$(GOPRIVATE) go mod tidy && go mod vendor

.PHONY: update
update:
	@GOPRIVATE=$(GOPRIVATE) go get -mod= -u all

.PHONY: install
install:
	@GOPRIVATE=$(GOPRIVATE) ROOT=$(dir $(BINPATH)) go generate ./tools.go
