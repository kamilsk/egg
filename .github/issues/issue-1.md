---
code: null
id: MDU6SXNzdWUyMDU1NTk5NzA=
databaseId: 205559970
number: 1
url: https://github.com/kamilsk/egg/issues/1
title: "add possibility to fix import via vanity urls"
issueType: null
labels:
  - "type:feature"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2017-02-06T11:50:14Z
updatedAt: 2017-04-22T19:09:32Z
lastEditedAt: 2017-02-08T18:49:53Z
closedAt: 2017-04-22T19:09:32Z
---

# add possibility to fix import via vanity urls

As example
```bash
$ go get bitbucket.org/kamilsk/semaphore | goget -fix-vanity-url
```
should be move `$GOPAHT/src/bitbucket.org/kamilsk/semaphore` to `$GOPAHT/github.com/kamilsk/semaphore`
