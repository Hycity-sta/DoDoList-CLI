package cmd

import "github.com/spf13/cobra"

func Lang() *cobra.Command {
	// 这里先保留语言命令的占位定义，后面再补实际逻辑。
	command := &cobra.Command{
		Use:   "lang [en|zh]",
		Short: "Set or show language.",
		Args:  cobra.RangeArgs(0, 1),
	}

	// 这里把语言命令的执行入口挂上去。
	command.Run = langRun
	return command
}

func langRun(cmd *cobra.Command, args []string) {
	// 语言切换逻辑暂时还没接入，这里先留空占位。
}
