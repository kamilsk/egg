---
id: 65
database_id: 1221893671
node_id: I_kwDOBNU1tM5I1J4n
status: open
title: "implement command to reduce versions of deps"
labels: []
url: https://github.com/kamilsk/egg/issues/65
created_at: 2022-04-30T18:34:14Z
updated_at: 2022-04-30T18:34:15Z
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
