package cmd

import (
	"dodolist/i18n"
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0"

func Version() *cobra.Command {
	// 这里只负责组装版本命令本身。
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdVersionUse),
		Short: i18n.T(i18n.CmdVersionShort),
		Run:   versionHandle,
	}
	return command
}

func versionHandle(cmd *cobra.Command, args []string) {
	// 版本命令只输出当前程序版本。
	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputVersion, version))
}
