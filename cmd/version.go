package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0"

func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information.",
		Run:   VersionHandle,
	}
}

func VersionHandle(cmd *cobra.Command, args []string) {
	fmt.Printf("dodolist %s\n", version)
}
