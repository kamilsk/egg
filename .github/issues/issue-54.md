---
code: null
id: MDU6SXNzdWU1OTI0MTQxNTY=
databaseId: 592414156
number: 54
url: https://github.com/kamilsk/egg/issues/54
title: "example mistake"
issueType: null
labels:
  - "type:bug"
  - "type:docs"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2020-04-02T07:22:07Z
updatedAt: 2020-04-02T19:53:16Z
lastEditedAt: null
closedAt: 2020-04-02T19:53:16Z
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
