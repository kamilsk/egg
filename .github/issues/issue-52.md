---
code: null
id: MDU6SXNzdWU1NDc2NzMxMDE=
databaseId: 547673101
number: 52
url: https://github.com/kamilsk/egg/issues/52
title: "invalid vendor mode detection"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2020-01-09T19:25:34Z
updatedAt: 2020-01-25T17:12:00Z
lastEditedAt: 2020-01-25T17:11:04Z
closedAt: 2020-01-25T17:12:00Z
---

# invalid vendor mode detection

if a parent dir has Gopkg.toml then gex define current manager as dep instead of current go mod
```
//go:generate go build -v -o=${ROOT}bin/mockgen ./vendor/github.com/golang/mock/mockgen
//go:generate go build -v -o=${ROOT}bin/golangci-lint ./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o=${ROOT}bin/egg ./vendor/github.com/kamilsk/egg
//go:generate go build -v -o=${ROOT}bin/easyjson ./vendor/github.com/mailru/easyjson/easyjson
//go:generate go build -v -o=${ROOT}bin/goimports ./vendor/golang.org/x/tools/cmd/goimports
```
- [x] reproduce problem
- [x] write smoke test
