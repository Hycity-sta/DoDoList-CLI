package cmd

import (
	"dodolist/storage"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Todo() *cobra.Command {
	var priority int

	command := &cobra.Command{
		Use:   "todo [content]",
		Short: "Create a new todo item.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validatePriority(priority); err != nil {
				return err
			}

			store := currentStore()
			// 创建时把内容、优先级和创建时间一次性写入。
			record := storage.Record{
				Note: storage.Note{
					Content:   joinContent(args),
					Priority:  priority,
					CreatedAt: time.Now(),
				},
			}

			if err := store.Append(record); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "created todo: %s\n", record.Content)
			return nil
		},
	}

	command.Flags().IntVar(&priority, "pro", 0, "priority of the todo item")
	return command
}
