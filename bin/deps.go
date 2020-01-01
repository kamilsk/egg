// +build tools

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	buf := new(bytes.Buffer)
	cmd := exec.CommandContext(ctx, "go", "mod", "edit", "-json")
	cmd.Stdout, cmd.Stderr = buf, os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var file struct {
		Module  Version
		Go      string
		Require []struct {
			Version
			Indirect bool // has "// indirect" comment
		}
		Exclude []Version
		Replace []struct {
			Old Version
			New Version
		}
	}
	if err := json.NewDecoder(buf).Decode(&file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	candidates := make([]interface{}, 0, len(file.Require))
	for _, req := range file.Require {
		if req.Indirect {
			continue
		}
		candidates = append(candidates, req.Path)
	}
	fmt.Println(candidates...)
}

// A Version is defined by a module path and version pair.
type Version struct {
	Path    string
	Version string `json:",omitempty"`
}
