package cmd

import "github.com/spf13/cobra"

func Help() *cobra.Command {
	// 这里定义帮助命令的基本信息。
	command := &cobra.Command{
		Use:   "help",
		Short: "Show help information.",
		Args:  cobra.NoArgs,
	}

	// 这里把帮助命令的执行逻辑挂上去。
	command.RunE = helpRunE
	return command
}

func helpRunE(cmd *cobra.Command, args []string) error {
	// 直接复用根命令自带的帮助输出。
	return cmd.Root().Help()
}
