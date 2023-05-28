---
code: null
id: MDU6SXNzdWU1NDYxMTkwNzE=
databaseId: 546119071
number: 51
url: https://github.com/kamilsk/egg/issues/51
title: "implement egg make build"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2020-01-07T07:09:36Z
updatedAt: 2021-01-12T12:52:37Z
lastEditedAt: null
closedAt: 2021-01-12T12:52:37Z
---

# implement egg make build

- [ ] `egg make build` as a port of https://github.com/octomation/makefiles/blob/452f5fe0bc6fba2828afaa58b9f254843a4d963a/dist/compiler.go
- [ ] use a map to prevent double inclusion of processed makefiles
- [ ] refactor to use a combination of working dir + absolute path of makefile

<!-- 2021-01-12T12:52Z https://github.com/kamilsk/egg/issues/51#issuecomment-758636004
won't do

now it's part of the [maintainer](https://github.com/octomation/maintainer) project.
-->
