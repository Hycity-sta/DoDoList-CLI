package cmd

import "github.com/spf13/cobra"

func Help() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Show help information.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().Help()
		},
	}
}
