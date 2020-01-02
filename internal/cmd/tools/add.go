package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewAddCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) error {
			tools, err := cfg.Create()
			if err != nil {
				return err
			}
			return tools.Add(context.TODO(), args...)
		},
	}
	return cmd
}
