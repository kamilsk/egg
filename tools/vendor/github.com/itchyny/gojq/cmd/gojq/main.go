package main

import (
	"os"

	"github.com/itchyny/gojq/cli"
)

func main() {
	os.Exit(cli.Run())
}
