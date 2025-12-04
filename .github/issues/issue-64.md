---
id: 64
database_id: 1217529302
node_id: I_kwDOBNU1tM5IkgXW
status: open
title: "security check"
labels: ["status:research","difficulty: medium"]
url: https://github.com/kamilsk/egg/issues/64
created_at: 2022-04-27T15:35:08Z
updated_at: 2022-04-27T15:35:09Z
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
