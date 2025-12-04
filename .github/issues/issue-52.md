---
id: 52
database_id: 547673101
node_id: MDU6SXNzdWU1NDc2NzMxMDE=
status: closed
title: "invalid vendor mode detection"
labels: ["type:bug"]
url: https://github.com/kamilsk/egg/issues/52
created_at: 2020-01-09T19:25:34Z
updated_at: 2020-01-25T17:12:00Z
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
