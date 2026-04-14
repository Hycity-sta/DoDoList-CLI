package cmd

import (
	"dodolist/utils"
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	var priority int
	var filterByPriority bool
	var sortByPriority bool
	var sortByStatus bool

	// 这里定义列表命令本身的基本信息。
	command := &cobra.Command{
		Use:   "list",
		Short: "List todo items.",
		Args:  cobra.NoArgs,
	}

	// 这里集中注册列表相关的筛选和排序参数。
	command.Flags().IntVar(&priority, "pro", 0, "filter todos by priority")
	command.Flags().BoolVar(&filterByPriority, "pro-filter", false, "enable priority filtering")
	command.Flags().BoolVar(&sortByPriority, "sort", false, "sort by priority descending")
	command.Flags().BoolVar(&sortByStatus, "status-sort", false, "sort by completion status")

	// 这个隐藏开关只作为内部状态承载，真实触发条件仍然是 --pro。
	_ = command.Flags().MarkHidden("pro-filter")

	// 这里分别挂上预处理逻辑和真正的执行逻辑。
	command.RunE = listRunE(&priority, &filterByPriority, &sortByPriority, &sortByStatus)
	command.PreRunE = listPreRunE(&priority, &filterByPriority)
	return command
}

func listRunE(priority *int, filterByPriority *bool, sortByPriority *bool, sortByStatus *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// 先读出当前所有待办，后续筛选和排序都基于这份数据。
		store := utils.CurrentStore()
		records, err := store.List()
		if err != nil {
			return err
		}

		// 先根据优先级过滤出需要展示的待办列表。
		items := utils.BuildViewItems(records)
		filtered := make([]utils.ViewItem, 0, len(items))
		for _, item := range items {
			// 只有显式传入 --pro 时才按优先级筛选。
			if *filterByPriority && item.Priority != *priority {
				continue
			}
			filtered = append(filtered, item)
		}

		// 先按优先级或创建时间完成基础排序。
		if *sortByPriority {
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

		// 如果要求按状态排序，就在基础顺序上再做一层稳定排序。
		if *sortByStatus {
			// 最后再按状态做稳定排序，这样状态会成为第一排序条件。
			sort.SliceStable(filtered, func(i, j int) bool {
				if filtered[i].Completed == filtered[j].Completed {
					return false
				}
				return !filtered[i].Completed && filtered[j].Completed
			})
		}

		// 最后把处理后的数据按表格形式输出到终端。
		// 用制表写出整齐表格，终端里更容易看。
		writer := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
		fmt.Fprintln(writer, "INDEX\tCREATED AT\tPRIORITY\tSTATUS\tTODO")
		for _, item := range filtered {
			fmt.Fprintf(writer, "%d\t%s\t%d\t%s\t%s\n", item.Index, utils.FormatTime(item.CreatedAt), item.Priority, utils.FormatStatus(item.Completed), item.Content)
		}
		return writer.Flush()
	}
}

func listPreRunE(priority *int, filterByPriority *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// 先根据是否传入 --pro 自动决定要不要开启优先级过滤。
		if cmd.Flags().Changed("pro") {
			*filterByPriority = true
		}

		// 如果开启了优先级过滤，就提前检查优先级是否合法。
		if *filterByPriority {
			return utils.ValidatePriority(*priority)
		}
		return nil
	}
}
