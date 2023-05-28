---
code: null
id: I_kwDOBNU1tM5IkgXW
databaseId: 1217529302
number: 64
url: https://github.com/kamilsk/egg/issues/64
title: "security check"
issueType: null
labels:
  - "status:research"
  - "difficulty: medium"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-04-27T15:35:08Z
updatedAt: 2022-04-27T15:35:09Z
lastEditedAt: null
closedAt: null
---

# security check

Add the possibility to overview all `go:generate` commands to prevent their execution by `go generate ./...` method.

See https://go.dev/blog/supply-chain. Make research: if it's a problem, then publish the issue to go. Try to figure out how go fixes it. Maybe `go generate` has some flags. Also, it will be great to show all `init()` functions.

```
$ egg scan
# scan go code and find
# - all go:generate directives for review
# - all init function
```
