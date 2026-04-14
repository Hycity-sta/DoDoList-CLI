package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Edit() *cobra.Command {
	var priority int
	var changePriority bool

	command := &cobra.Command{
		Use:   "edit [index] [content]",
		Short: "Edit a todo item.",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := parseIndexArg(args[0])
			if err != nil {
				return err
			}
			if changePriority {
				if err := validatePriority(priority); err != nil {
					return err
				}
			}

			store := currentStore()
			records, err := mustRecords(store)
			if err != nil {
				return err
			}
			if index >= len(records) {
				return fmt.Errorf("todo %d does not exist", index+1)
			}

			// 编辑时始终更新内容，优先级只有传了 --pro 才改。
			records[index].Content = joinContent(args[1:])
			if changePriority {
				records[index].Priority = priority
			}

			if err := store.Replace(records); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "edited todo %d\n", index+1)
			return nil
		},
	}

	command.Flags().IntVar(&priority, "pro", 0, "new priority of the todo item")
	command.PreRunE = func(cmd *cobra.Command, args []string) error {
		changePriority = cmd.Flags().Changed("pro")
		return nil
	}
	return command
}
