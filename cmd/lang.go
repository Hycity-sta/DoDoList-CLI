package cmd

import (
	"dodolist/config"
	"dodolist/i18n"
	"fmt"

	"github.com/spf13/cobra"
)

func Lang() *cobra.Command {
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdLangUse),
		Short: i18n.T(i18n.CmdLangShort),
		Args:  cobra.RangeArgs(0, 1),
	}

	command.RunE = langHandle
	return command
}

func langHandle(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.CmdLangCurrent, config.Language()))
		return err
	}

	lang := args[0]
	if !config.IsSupportedLanguage(lang) {
		return fmt.Errorf(i18n.T(i18n.CmdLangUnknown, lang))
	}

	if err := config.SetLanguage(lang); err != nil {
		return err
	}
	i18n.SetLang(config.Language())
	_, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.CmdLangSet, i18n.Lang()))
	return err
}
