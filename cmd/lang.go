package cmd

import "github.com/spf13/cobra"

func Lang() *cobra.Command {
	return &cobra.Command{
		Use:   "lang [en|zh]",
		Short: "Set or show language.",
		Args:  cobra.RangeArgs(0, 1),
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
