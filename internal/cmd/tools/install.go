package tools

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewInstallCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "install",
		RunE: func(_ *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			if _, present := os.LookupEnv("GOBIN"); !present {
				return errors.New("GOBIN is not set")
			}
			if _, err := os.Stat(cfg.ManifestName); err != nil {
				return err
			}

			return exec.CommandContext(ctx, "go", "generate", cfg.ManifestName).Run()
		},
	}
	return cmd
}
