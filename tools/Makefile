.DEFAULT_GOAL = install

SHELL = /bin/bash -euo pipefail

BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   =
GOPROXY     = direct
GOTAGS      = -tags tools
MODULE      = $(shell go list -m)

export PATH := $(BINPATH):$(PATH)

.PHONY: tools-env
tools-env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell go env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell go env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell go env GOPROXY))"
	@echo "GOTAGS:      $(GOTAGS)"
	@echo "MODULE:      $(MODULE)"

.PHONY: deps
deps:
	@GOPRIVATE=$(GOPRIVATE) go mod tidy && go mod vendor

.PHONY: update
update:
	@GOPRIVATE=$(GOPRIVATE) go get -tags tools -mod= -u

.PHONY: install
install:
	@GOPRIVATE=$(GOPRIVATE) ROOT=$(dir $(BINPATH)) go generate ./tools.go


.PHONY: env
env: tools-env
