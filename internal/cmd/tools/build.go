package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			tools, err := new(gex.Config).Create()
			if err != nil {
				return err
			}
			return tools.BuildAll(context.TODO())
		},
	}
	return cmd
}
