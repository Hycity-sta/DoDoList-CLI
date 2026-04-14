package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Delete() *cobra.Command {
	return &cobra.Command{
		Use:     "ok [index]",
		Aliases: []string{"delete"},
		Short:   "Complete a todo item.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := parseIndexArg(args[0])
			if err != nil {
				return err
			}

			store := currentStore()
			records, err := mustRecords(store)
			if err != nil {
				return err
			}
			if index >= len(records) {
				return fmt.Errorf("todo %d does not exist", index+1)
			}

			// 完成待办的实现就是把对应项从存储切片里移除。
			records = append(records[:index], records[index+1:]...)
			if err := store.Replace(records); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "completed todo %d\n", index+1)
			return nil
		},
	}
}
