package cmd

import (
	"dodolist/i18n"

	"github.com/spf13/cobra"
)

func Help() *cobra.Command {
	// 这里定义帮助命令的基本信息。
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdHelpUse),
		Short: i18n.T(i18n.CmdHelpShort),
		Args:  cobra.NoArgs,
	}

	// 这里把帮助命令的执行逻辑挂上去。
	command.RunE = helpHandle
	return command
}

func helpHandle(cmd *cobra.Command, args []string) error {
	// 直接复用根命令自带的帮助输出。
	return cmd.Root().Help()
}
