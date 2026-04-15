package cmd

import (
	"dodolist/i18n"
	"dodolist/storage"
	"dodolist/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// 用于清理所有已完成的待办
func Clear() *cobra.Command {
	// 这里定义清理已完成待办命令的基本信息
	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdClearUse),
		Short: i18n.T(i18n.CmdClearShort),
		Args:  cobra.NoArgs,
	}

	// 这里把清理逻辑挂到命令对象上
	command.RunE = clearHandle
	return command
}

func clearHandle(cmd *cobra.Command, args []string) error {
	// 先读取当前全部待办，再筛出仍需保留的未完成项
	store := utils.CurrentStore()
	records, err := store.List()
	if err != nil {
		return err
	}

	active := make([]storage.Record, 0, len(records))
	cleared := 0
	for _, record := range records {
		// 已完成的待办直接计数并跳过，其他待办继续保留
		if record.Completed {
			cleared++
			continue
		}
		active = append(active, record)
	}

	// 最后把清理后的结果整体写回文件并输出提示
	if err := store.Replace(active); err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputClearedTodo, cleared))
	return nil
}
