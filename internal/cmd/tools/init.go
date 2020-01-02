package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewInitCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		RunE: func(*cobra.Command, []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			repository, err := cfg.Create()
			if err != nil {
				return err
			}
			return repository.Add(ctx)
		},
	}
	return cmd
}
