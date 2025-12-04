---
id: 51
database_id: 546119071
node_id: MDU6SXNzdWU1NDYxMTkwNzE=
status: closed
title: "implement egg make build"
labels: ["type:task"]
url: https://github.com/kamilsk/egg/issues/51
created_at: 2020-01-07T07:09:36Z
updated_at: 2021-01-12T12:52:37Z
---

# implement egg make build

- [ ] `egg make build` as a port of https://github.com/octomation/makefiles/blob/452f5fe0bc6fba2828afaa58b9f254843a4d963a/dist/compiler.go
- [ ] use a map to prevent double inclusion of processed makefiles
- [ ] refactor to use a combination of working dir + absolute path of makefile
