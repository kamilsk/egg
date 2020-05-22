> # 🐣 egg
>
> Extended go get - alternative for the standard `go get` with a few little but useful features.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

```bash
$ egg tools  add golang.org/x/tools/cmd/goimports@latest

$ egg binary get github.com/goreleaser/goreleaser@latest
```

A full description of the idea is available [here][design.page].

## 🏆 Motivation

At [Avito](https://tech.avito.ru/), we develop many services written on [Go](https://golang.org),
and many of them use different tools such as [easyjson](https://github.com/mailru/easyjson),
[goimports](https://github.com/kamilsk/go-tools/releases/tag/goimports),
[mockgen](https://github.com/golang/mock), and others. It isn't effortless to manage all of them
through all teams in the company.

I need a tool that helps me to track tool dependencies for services and interact with them.

## 🤼‍♂️ How to

### egg deps

```bash
$ egg deps list
> github.com/izumin5210/gex
> github.com/spf13/cobra
> github.com/stretchr/testify
> go.octolab.org/toolkit/cli

$ egg deps check ...
```

### egg tools

```bash
$ mkdir tools && cd tools

# init a new toolset
$ go mod init your.module/tools
$ egg tools init

# add golangci-lint to tools.go
$ egg tools add github.com/golangci/golangci-lint/cmd/golangci-lint
# add mockgen to tools.go and build it to bin/mockgen
$ egg tools add --build github.com/golang/mock/mockgen

# install tools to ${GOPATH}/bin
$ export GOBIN="${GOPATH}/bin"
$ go generate tools.go
# do the same thing
$ egg tools install

# list the toolset
$ egg tools list
> mockgen
> golangci-lint
```

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/egg
```

### Binary

```bash
$ curl -sSfL https://raw.githubusercontent.com/kamilsk/egg/master/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/kamilsk/egg/master/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/kamilsk/egg@latest
# or use egg tool
$ egg tools add github.com/kamilsk/egg@latest

$ go mod edit -replace=github.com/izumin5210/gex=github.com/kamilsk/gex@latest
```

### Bash and Zsh completions

```bash
$ egg completion bash > /path/to/bash_completion.d/egg.sh
$ egg completion zsh  > /path/to/zsh-completions/_egg.zsh
# or autodetect
$ source <(egg completion)
```

> See `kubectl` [documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion).

## 🤲 Outcomes

### Patches

- [github.com/izumin5210/gex](https://github.com/izumin5210/gex)
  - [differences from the upstream](https://github.com/izumin5210/gex/compare/master...kamilsk:extended)

### Research

- [Proposals][rfc.page]
- [Related projects and articles][research.page]

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/kamilsk/egg
[build.icon]:       https://travis-ci.com/kamilsk/egg.svg?branch=master
[design.page]:      https://www.notion.so/octolab/egg-f716b80d4b184301b1db2e058f603dd0?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/kamilsk/egg
[research.page]:    https://github.com/under-the-hood/egg
[rfc.page]:         https://github.com/octolab/rfc/pulls?utf8=✓&q=is%3Apr+label%3Akamilsk%2Fegg+
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue
