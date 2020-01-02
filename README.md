> # 🐣 egg
>
> Extended go get - alternative for the standard "go get" with a few little but useful features.

[![Build][build.icon]][build.page]

## 💡 Idea

...

## 🏆 Motivation

...

## 🤼‍♂️ How to

### egg deps

```bash
$ egg deps ...
```

### egg make

```bash
$ egg make ...
```

### egg tools

```bash
$ mkdir tools && cd tools
$ go mod init your.module/tools
$ egg tools init
# add golangci-lint to your toolset
$ egg tools add github.com/golangci/golangci-lint
# add mockgen to your toolset and build it to bin/mockgen
$ egg tools add --build github.com/golang/mock/mockgen
# build tools to ${GOPATH}/bin/
$ ROOT="${GOPATH}/" go generate tools.go
```

### egg vanity

```bash
$ egg vanity ...
```

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/egg
```

### Binary

```bash
$ curl -sSL https://bit.ly/-chicken- | sh
# or
$ wget -qO- https://bit.ly/-chicken- | sh
```

### Source

```bash
$ go get -u github.com/kamilsk/egg
# or
$ egg tools add github.com/kamilsk/egg
```

### Bash and Zsh completions

```bash
$ egg completion bash > /path/to/bash_completion.d/egg.sh
$ egg completion zsh  > /path/to/zsh-completions/_egg.zsh
```

## 🤲 Outcomes

### Patches

- [github.com/izumin5210/gex](https://github.com/izumin5210/gex)
  - [fork](https://github.com/izumin5210/gex/compare/master...kamilsk:extended)
  - `replace github.com/izumin5210/gex => github.com/kamilsk/gex latest`

---

made with ❤️ for everyone

[build.icon]:       https://travis-ci.org/kamilsk/egg.svg?branch=master
[build.page]:       https://travis-ci.org/kamilsk/egg

[promo.page]:       https://github.com/kamilsk/egg
