package cmd

import (
	"dodolist/i18n"
	"dodolist/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func Ok() *cobra.Command {
	// 这里定义“完成待办”命令的基本信息。
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdOkUse),
		Short: i18n.T(i18n.CmdOkShort),
		Args:  cobra.ExactArgs(1),
	}

	// 这里把命令逻辑挂到 Handle 上。
	command.RunE = okHandle
	return command
}

func okHandle(cmd *cobra.Command, args []string) error {
	// 先把命令行索引转换成内部切片下标。
	index, err := utils.ParseIndex(args[0])
	if err != nil {
		return err
	}

	// 再从存储中读取当前全部待办数据。
	store := utils.CurrentStore()
	records, err := store.List()
	if err != nil {
		return err
	}
	if index >= len(records) {
		return fmt.Errorf(i18n.T(i18n.ErrTodoNotExist, index+1))
	}

	// 把目标待办标记为已完成。
	records[index].Completed = true
	if err := store.Replace(records); err != nil {
		return err
	}

	// 最后输出完成提示。
	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputCompletedTodo, index+1))
	return nil
}
