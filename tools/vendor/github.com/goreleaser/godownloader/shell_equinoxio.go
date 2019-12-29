package main

import (
	"fmt"
	"path"

	"github.com/goreleaser/goreleaser/pkg/config"
)

// processEquinoxio create a fake goreleaser config for equinox.io
// and use a similar template.
func processEquinoxio(repo string) ([]byte, error) {
	if repo == "" {
		return nil, fmt.Errorf("must have repo")
	}
	project := config.Project{}
	project.Release.GitHub.Owner = path.Dir(repo)
	project.Release.GitHub.Name = path.Base(repo)
	project.Builds = []config.Build{
		{Binary: path.Base(repo)},
	}
	project.Archive.Format = "tgz"
	return makeShell(shellEquinoxio, &project)
}

const shellEquinoxio = `#!/bin/sh
set -e
# Code generated by godownloader on {{ timestamp }}. DO NOT EDIT.
#

usage() {
  this=$1
  cat <<EOF

$this: download go binaries for {{ $.Release.GitHub.Owner }}/{{ $.Release.GitHub.Name }}

Usage: $this [-b bindir]
  -b set the installation or BINDIR, defaults to ./bin
  -d turns on debug logging

  Version is automatically selected from
  https://dl.equinox.io/{{ $.Release.GitHub.Owner }}/{{ $.Release.GitHub.Name }}/stable

Generated by godownloader
 https://github.com/goreleaser/godownloader

EOF
  exit 2
}

parse_args() {
  #BINDIR is ./bin unless set be ENV
  # over-ridden by flag below

  BINDIR=${BINDIR:-./bin}
  while getopts "b:dh?x" arg; do
    case "$arg" in
      b) BINDIR="$OPTARG" ;;
      d) log_set_priority 10 ;;
      h | \?) usage "$0" ;;
      x) set -x ;;
    esac
  done
  shift $((OPTIND - 1))
  # VERSION currently unused
  #VERSION=$1
}
# wrap all destructive operations into a function
# to prevent curl|bash network truncation and disaster
execute() {
  TMPDIR=$(mktemp -d)

  log_info "seeking '${CHANNEL}' latest from $TARGET"
  TARBALL_URL=$(http_copy "$TARGET" | grep "$TARBALL" | cut -d '"' -f 2)

  log_info "downloading from ${TARBALL_URL}"
  http_download "${TMPDIR}/${TARBALL}" "$TARBALL_URL"

  (cd "$TMPDIR" && untar "$TARBALL")
  test ! -d "${BINDIR}" && install -d "${BINDIR}"
  install "${TMPDIR}/${BINARY}" "${BINDIR}/"
  log_info "installed ${BINDIR}/${BINARY}"
}` + shellfn + `OWNER={{ .Release.GitHub.Owner }}
REPO="{{ .Release.GitHub.Name }}"
BINARY={{ (index .Builds 0).Binary }}
FORMAT={{ .Archive.Format }}
BINDIR=${BINDIR:-./bin}
CHANNEL=stable
PREFIX="$OWNER/$REPO"
# use in logging routines
log_prefix() {
        echo "$PREFIX"
}
OS=$(uname_os)
ARCH=$(uname_arch)

uname_os_check "$OS"
uname_arch_check "$ARCH"

parse_args "$@"

TARGET=https://dl.equinox.io/${OWNER}/${REPO}/${CHANNEL}
TARBALL="${BINARY}-${CHANNEL}-${OS}-${ARCH}.${FORMAT}"

if [ "$OS" = "windows" ]; then
  BINARY="${BINARY}.exe"
fi

execute
`
