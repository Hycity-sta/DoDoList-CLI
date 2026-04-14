package main

import (
	"dodolist/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// 在入口处统一定义根命令和所有子命令。
	root := &cobra.Command{
		Use:   "dodolist",
		Short: "DoDoList is a CLI tool for managing todo items.",
		Long:  "DoDoList stores todo items in a single local JSON file.",
	}

	// 这里直接展开注册，避免再绕一层组装函数。
	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Help())
	root.AddCommand(cmd.Lang())
	root.AddCommand(cmd.Todo())
	root.AddCommand(cmd.List())
	root.AddCommand(cmd.Delete())
	root.AddCommand(cmd.Edit())

	// Cobra 返回错误时直接以非零状态退出，方便脚本判断执行结果。
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
