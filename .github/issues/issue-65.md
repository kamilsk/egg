---
code: null
id: I_kwDOBNU1tM5I1J4n
databaseId: 1221893671
number: 65
url: https://github.com/kamilsk/egg/issues/65
title: "implement command to reduce versions of deps"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-04-30T18:34:14Z
updatedAt: 2022-04-30T18:34:15Z
lastEditedAt: null
closedAt: null
---

# implement command to reduce versions of deps

```bash
$ egg deps reduce --go=1.11
```

the algorithm:

- run test and check coverage
- define list of deps with downgrade versions
- filter them by `--go`
- if nothing available - return error
- otherwise use binary search to define optimal min version
- run test everytime a dep was updated
