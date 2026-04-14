package cmd

import (
	"dodolist/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func Ok() *cobra.Command {
	// 这里定义“完成待办”命令的基本信息。
	command := &cobra.Command{
		Use:   "ok [index]",
		Short: "Mark a todo item as completed.",
		Args:  cobra.ExactArgs(1),
	}

	// 这里把命令逻辑挂到 RunE 上。
	command.RunE = okRunE
	return command
}

func okRunE(cmd *cobra.Command, args []string) error {
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
		return fmt.Errorf("todo %d does not exist", index+1)
	}

	// 把目标待办标记为已完成。
	records[index].Completed = true
	if err := store.Replace(records); err != nil {
		return err
	}

	// 最后输出完成提示。
	fmt.Fprintf(cmd.OutOrStdout(), "completed todo %d\n", index+1)
	return nil
}
