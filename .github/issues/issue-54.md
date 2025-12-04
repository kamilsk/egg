---
id: 54
database_id: 592414156
node_id: MDU6SXNzdWU1OTI0MTQxNTY=
status: closed
title: "example mistake"
labels: ["type:bug","type:docs"]
url: https://github.com/kamilsk/egg/issues/54
created_at: 2020-04-02T07:22:07Z
updated_at: 2020-04-02T19:53:16Z
---

# example mistake

instead of

```
$ egg tools add github.com/golangci/golangci-lint
```

must be

```
$ egg tools add github.com/golangci/golangci-lint/cmd/golangci-lint
```
