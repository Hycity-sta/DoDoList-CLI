package cmd

import "github.com/spf13/cobra"

var storePathOverride string

func RootForTest(path string) *cobra.Command {
	// 测试时把数据文件重定向到临时目录，避免污染真实待办数据。
	storePathOverride = path
	root := &cobra.Command{
		Use:   "dodolist",
		Short: "DoDoList is a CLI tool for managing todo items.",
		Long:  "DoDoList stores todo items in a single local JSON file.",
	}
	root.AddCommand(Version())
	root.AddCommand(Help())
	root.AddCommand(Lang())
	root.AddCommand(Todo())
	root.AddCommand(List())
	root.AddCommand(Delete())
	root.AddCommand(Edit())
	return root
}
