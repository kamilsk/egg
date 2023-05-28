---
code: null
id: MDU6SXNzdWUyMjQ2ODI1NzI=
databaseId: 224682572
number: 14
url: https://github.com/kamilsk/egg/issues/14
title: "fix readme"
issueType: null
labels:
  - "type:bug"
  - "type:docs"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2017-04-27T06:24:44Z
updatedAt: 2017-04-28T09:19:00Z
lastEditedAt: null
closedAt: 2017-04-28T09:19:00Z
---

# fix readme

`go get [package with vanity url problem] 2>&1 | egg -fix-vanity-url -version 2.x`

`-version 2.x` - impossible in this case, because nothing happens in stdout/stderr

```
$ go get golang.org/x/net/context 2>&1
...nothing
```
