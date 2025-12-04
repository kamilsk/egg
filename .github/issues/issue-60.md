---
id: 60
database_id: 784210191
node_id: MDU6SXNzdWU3ODQyMTAxOTE=
status: open
title: "egg deps list not consistent"
labels: ["type:bug"]
url: https://github.com/kamilsk/egg/issues/60
created_at: 2021-01-12T12:54:51Z
updated_at: 2021-01-12T12:54:51Z
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
