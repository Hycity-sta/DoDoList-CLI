package cmd

import (
	"dodolist/storage"
	"dodolist/utils"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Todo() *cobra.Command {
	var priority int

	// 这里定义创建待办命令的基本信息。
	command := &cobra.Command{
		Use:   "todo [content]",
		Short: "Create a new todo item.",
		Args:  cobra.MinimumNArgs(1),
	}

	// 这里挂载命令所需的标记和执行入口。
	command.Flags().IntVar(&priority, "pro", 0, "priority of the todo item")
	command.RunE = todoRunE(&priority)
	return command
}

func todoRunE(priority *int) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// 先校验优先级，避免把非法值写进存储。
		if err := utils.ValidatePriority(*priority); err != nil {
			return err
		}

		// 先拿到当前存储对象，后面统一走单文件持久化。
		store := utils.CurrentStore()

		// 创建时把内容、优先级和创建时间一次性写入。
		record := storage.Record{
			Note: storage.Note{
				Content:   utils.JoinContent(args),
				Priority:  *priority,
				CreatedAt: time.Now(),
			},
		}

		// 把新待办追加到文件里。
		if err := store.Append(record); err != nil {
			return err
		}

		// 最后把创建结果打印给终端用户。
		fmt.Fprintf(cmd.OutOrStdout(), "created todo: %s\n", record.Content)
		return nil
	}
}
