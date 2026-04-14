package cmd

import (
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	var priority int
	var filterByPriority bool
	var sortByPriority bool

	command := &cobra.Command{
		Use:   "list",
		Short: "List todo items.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			store := currentStore()
			records, err := mustRecords(store)
			if err != nil {
				return err
			}

			items := recordView(records)
			filtered := make([]todoView, 0, len(items))
			for _, item := range items {
				// 只有显式传入 --pro 时才按优先级筛选。
				if filterByPriority && item.Priority != priority {
					continue
				}
				filtered = append(filtered, item)
			}

			if sortByPriority {
				// 按优先级从高到低排，相同优先级再按创建时间排。
				sort.SliceStable(filtered, func(i, j int) bool {
					if filtered[i].Priority == filtered[j].Priority {
						return filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
					}
					return filtered[i].Priority > filtered[j].Priority
				})
			} else {
				// 默认保持按创建时间升序展示。
				sort.SliceStable(filtered, func(i, j int) bool {
					return filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
				})
			}

			// 用制表写出整齐表格，终端里更容易看。
			writer := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(writer, "INDEX\tCREATED AT\tPRIORITY\tTODO")
			for _, item := range filtered {
				fmt.Fprintf(writer, "%d\t%s\t%d\t%s\n", item.Index, formatTime(item.CreatedAt), item.Priority, item.Content)
			}
			return writer.Flush()
		},
	}

	command.Flags().IntVar(&priority, "pro", 0, "filter todos by priority")
	command.Flags().BoolVar(&filterByPriority, "pro-filter", false, "enable priority filtering")
	command.Flags().BoolVar(&sortByPriority, "sort", false, "sort by priority descending")
	// 这个隐藏开关只作为内部状态承载，真实触发条件仍然是 --pro。
	_ = command.Flags().MarkHidden("pro-filter")
	command.PreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("pro") {
			filterByPriority = true
		}
		if filterByPriority {
			return validatePriority(priority)
		}
		return nil
	}
	return command
}
