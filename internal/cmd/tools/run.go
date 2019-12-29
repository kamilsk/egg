package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			tools, err := new(gex.Config).Create()
			if err != nil {
				return err
			}
			return tools.Run(context.TODO(), args[0], args[1:]...)
		},
	}
	return cmd
}
