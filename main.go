package main

import (
	"dodolist/cmd"
	"dodolist/config"
	"dodolist/i18n"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	i18n.SetLang(config.Language())

	// 在入口处统一定义根命令和所有子命令。
	root := &cobra.Command{
		Use:   "dodolist",
		Short: i18n.T(i18n.AppShort),
		Long:  i18n.T(i18n.AppLong),
	}

	// 这里直接展开注册，避免再绕一层组装函数。
	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Help())
	root.AddCommand(cmd.Lang())
	root.AddCommand(cmd.Todo())
	root.AddCommand(cmd.Show())
	root.AddCommand(cmd.Ok())
	root.AddCommand(cmd.Delete())
	root.AddCommand(cmd.Edit())

	// Cobra 返回错误时直接以非零状态退出，方便脚本判断执行结果。
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
