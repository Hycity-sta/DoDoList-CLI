package cmd

import (
	"dodolist/i18n"
	"dodolist/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func Edit() *cobra.Command {
	var priority int
	var changePriority bool

	// 这里定义编辑命令本身的基本信息。
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdEditUse),
		Short: i18n.T(i18n.CmdEditShort),
		Args:  cobra.MinimumNArgs(2),
	}

	// 这里注册编辑命令的参数和执行流程。
	command.Flags().IntVar(&priority, "pro", 0, i18n.T(i18n.CmdEditPriority))
	command.RunE = editHandle(&priority, &changePriority)
	command.PreRunE = editPreHandle(&changePriority)
	return command
}

func editHandle(priority *int, changePriority *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// 先把命令行索引转换成内部切片下标。
		index, err := utils.ParseIndex(args[0])
		if err != nil {
			return err
		}

		// 如果这次要改优先级，就先做优先级校验。
		if *changePriority {
			if err := utils.ValidatePriority(*priority); err != nil {
				return err
			}
		}

		// 再读取现有待办，准备在内存里修改目标项。
		store := utils.CurrentStore()
		records, err := store.List()
		if err != nil {
			return err
		}
		if index >= len(records) {
			return fmt.Errorf(i18n.T(i18n.ErrTodoNotExist, index+1))
		}

		// 这里按输入更新待办内容和可选的优先级。
		// 编辑时始终更新内容，优先级只有传了 --pro 才改。
		records[index].Content = utils.JoinContent(args[1:])
		if *changePriority {
			records[index].Priority = *priority
		}

		// 修改完成后把整份数据重新写回文件。
		if err := store.Replace(records); err != nil {
			return err
		}

		// 最后输出编辑成功提示。
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputEditedTodo, index+1))
		return nil
	}
}

func editPreHandle(changePriority *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// 根据是否传入 --pro 决定这次是否需要修改优先级。
		*changePriority = cmd.Flags().Changed("pro")
		return nil
	}
}
