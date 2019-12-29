package tools

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools",
		Short: "Manage tools",
		Long:  "Manage tools.",
	}
	cmd.AddCommand(
		NewAddCommand(),
		NewBuildCommand(),
		NewInitCommand(),
		NewRegenCommand(),
		NewRunCommand(),
	)
	return cmd
}
