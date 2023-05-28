---
code: null
id: MDU6SXNzdWU3ODQyMTAxOTE=
databaseId: 784210191
number: 60
url: https://github.com/kamilsk/egg/issues/60
title: "egg deps list not consistent"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2021-01-12T12:54:51Z
updatedAt: 2021-01-12T12:54:51Z
lastEditedAt: null
closedAt: null
---

# egg deps list not consistent

```bash
$ go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m -mod=readonly all
github.com/golang/mock
github.com/spf13/cobra
github.com/stretchr/testify
go.octolab.org
go.octolab.org/toolkit/cli
go.octolab.org/toolkit/config
```

```bash
$ egg deps list
github.com/golang/mock github.com/spf13/cobra github.com/stretchr/testify go.octolab.org go.octolab.org/toolkit/cli go.octolab.org/toolkit/config
```
