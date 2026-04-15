package cmd

import (
	"dodolist/i18n"
	"dodolist/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// 用于删除待办事项
func Delete() *cobra.Command {
	// 这里定义删除命令的基本信息
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdDeleteUse),
		Short: i18n.T(i18n.CmdDeleteShort),
		Args:  cobra.ExactArgs(1),
	}

	// 这里把删除逻辑挂到命令对象上
	command.RunE = deleteHandle
	return command
}

func deleteHandle(cmd *cobra.Command, args []string) error {
	// 先把命令行索引转换成内部切片下标
	index, err := utils.ParseIndex(args[0])
	if err != nil {
		return err
	}

	// 再把现有待办列表从存储里读出来
	store := utils.CurrentStore()
	records, err := store.List()
	if err != nil {
		return err
	}
	if index >= len(records) {
		return fmt.Errorf(i18n.T(i18n.ErrTodoNotExist, index+1))
	}

	// 从切片里移除目标待办，再整体回写到文件
	// 删除待办的实现就是把对应项从存储切片里移除
	records = append(records[:index], records[index+1:]...)
	if err := store.Replace(records); err != nil {
		return err
	}

	// 最后把删除结果打印出来
	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputDeletedTodo, index+1))
	return nil
}
