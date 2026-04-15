package cmd

import (
	"dodolist/config"
	"dodolist/i18n"
	"fmt"

	"github.com/spf13/cobra"
)

// 用于切换语言
func Lang() *cobra.Command {
	// 这里定义语言命令的基本信息
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdLangUse),
		Short: i18n.T(i18n.CmdLangShort),
		Args:  cobra.RangeArgs(0, 1),
	}

	// 这里挂上语言查看和切换逻辑
	command.RunE = langHandle
	return command
}

func langHandle(cmd *cobra.Command, args []string) error {
	// 不传参数时直接显示当前语言配置
	if len(args) == 0 {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.CmdLangCurrent, config.Language()))
		return err
	}

	// 传入语言值时，先校验是否属于支持列表
	lang := args[0]
	if !config.IsSupportedLanguage(lang) {
		return fmt.Errorf(i18n.T(i18n.CmdLangUnknown, lang))
	}

	// 校验通过后写入配置，并同步刷新当前 i18n 状态
	if err := config.SetLanguage(lang); err != nil {
		return err
	}
	i18n.SetLang(config.Language())
	_, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.CmdLangSet, i18n.Lang()))
	return err
}
