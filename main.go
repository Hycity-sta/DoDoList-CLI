package main

import (
	"dodolist/cmd"
	"dodolist/config"
	"dodolist/i18n"
	"fmt"
	"os"
)

func main() {
	// 加载配置文件
	if err := config.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 读取配置文件设置语言
	i18n.SetLang(config.Language())

	// 获取所有命令
	cmdline := cmd.Setup()

	// Cobra 返回错误时直接以非零状态退出，方便脚本判断执行结果。
	if err := cmdline.Execute(); err != nil {
		os.Exit(1)
	}
}
