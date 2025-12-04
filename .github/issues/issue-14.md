---
id: 14
database_id: 224682572
node_id: MDU6SXNzdWUyMjQ2ODI1NzI=
status: closed
title: "fix readme"
labels: ["type:bug","type:docs"]
url: https://github.com/kamilsk/egg/issues/14
created_at: 2017-04-27T06:24:44Z
updated_at: 2017-04-28T09:19:00Z
---

# fix readme

`go get [package with vanity url problem] 2>&1 | egg -fix-vanity-url -version 2.x`

`-version 2.x` - impossible in this case, because nothing happens in stdout/stderr

```
$ go get golang.org/x/net/context 2>&1
...nothing
```
