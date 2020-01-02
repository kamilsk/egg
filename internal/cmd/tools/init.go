package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

var defaults = []string{
	"github.com/golang/mock/mockgen",
	"github.com/golangci/golangci-lint/cmd/golangci-lint",
	"golang.org/x/tools/cmd/goimports",
}

func NewInitCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		RunE: func(*cobra.Command, []string) error {
			tools, err := cfg.Create()
			if err != nil {
				return err
			}
			return tools.Add(context.TODO(), defaults...)
		},
	}
	return cmd
}
