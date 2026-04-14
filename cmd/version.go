package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0"

func Version() *cobra.Command {
	// 这里只负责组装版本命令本身。
	command := &cobra.Command{
		Use:   "version",
		Short: "Print version information.",
		Run:   versionHandle,
	}
	return command
}

func versionHandle(cmd *cobra.Command, args []string) {
	// 版本命令只输出当前程序版本。
	fmt.Printf("dodolist %s\n", version)
}
