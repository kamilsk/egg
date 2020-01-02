package tools

import (
	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools",
		Short: "Manage tools",
		Long:  "Manage tools.",
	}
	cfg := &gex.Config{
		OutWriter: cmd.OutOrStdout(),
		ErrWriter: cmd.ErrOrStderr(),
		InReader:  cmd.InOrStdin(),
	}
	cmd.AddCommand(
		NewAddCommand(cfg),
		NewBuildCommand(cfg),
		NewInitCommand(cfg),
		NewRegenCommand(cfg),
		NewRunCommand(cfg),
	)
	return cmd
}
