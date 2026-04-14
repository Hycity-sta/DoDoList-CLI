package cmd

import (
	"dodolist/i18n"
	"dodolist/utils"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func Show() *cobra.Command {
	var priority int
	var filterByPriority bool
	var sortByPriority bool
	var sortByStatus bool

	// 这里定义 show 命令本身的基本信息。
	command := &cobra.Command{
		Use:     i18n.T(i18n.CmdShowUse),
		Short:   i18n.T(i18n.CmdShowShort),
		Args:    cobra.NoArgs,
		Aliases: []string{"list"},
	}

	// 这里集中注册 show 相关的筛选和排序参数。
	command.Flags().IntVar(&priority, "pro", 0, i18n.T(i18n.CmdShowPriority))
	command.Flags().BoolVar(&filterByPriority, "pro-filter", false, i18n.T(i18n.CmdShowProFilter))
	command.Flags().BoolVar(&sortByPriority, "sort", false, i18n.T(i18n.CmdShowSort))
	command.Flags().BoolVar(&sortByStatus, "status-sort", false, i18n.T(i18n.CmdShowStatusSort))

	// 这个隐藏开关只作为内部状态承载，真实触发条件仍然是 --pro。
	_ = command.Flags().MarkHidden("pro-filter")

	// 这里分别挂上预处理逻辑和真正的执行逻辑。
	command.RunE = showHandle(&priority, &filterByPriority, &sortByPriority, &sortByStatus)
	command.PreRunE = showPreHandle(&priority, &filterByPriority)
	return command
}

func showHandle(priority *int, filterByPriority *bool, sortByPriority *bool, sortByStatus *bool) func(cmd *cobra.Command, args []string) error {
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
		// 先按终端可见宽度计算每一列的最小宽度，避免中文字符把表头挤歪。
		headers := []string{
			i18n.T(i18n.CmdShowHeaderIndex),
			i18n.T(i18n.CmdShowHeaderCreatedAt),
			i18n.T(i18n.CmdShowHeaderPriority),
			i18n.T(i18n.CmdShowHeaderStatus),
			i18n.T(i18n.CmdShowHeaderTodo),
		}
		rows := make([][]string, 0, len(filtered))
		for _, item := range filtered {
			rows = append(rows, []string{
				fmt.Sprintf("%d", item.Index),
				utils.FormatTime(item.CreatedAt),
				fmt.Sprintf("%d", item.Priority),
				utils.FormatStatus(item.Completed),
				item.Content,
			})
		}

		widths := make([]int, len(headers))
		for i, header := range headers {
			widths[i] = utils.DisplayWidth(header)
		}
		for _, row := range rows {
			for i, cell := range row {
				if width := utils.DisplayWidth(cell); width > widths[i] {
					widths[i] = width
				}
			}
		}

		writeRow := func(values []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s  %s  %s  %s  %s\n",
				utils.PadRightDisplay(values[0], widths[0]),
				utils.PadRightDisplay(values[1], widths[1]),
				utils.PadRightDisplay(values[2], widths[2]),
				utils.PadRightDisplay(values[3], widths[3]),
				utils.PadRightDisplay(values[4], widths[4]),
			)
		}

		writeRow(headers)
		for _, row := range rows {
			writeRow(row)
		}
		return nil
	}
}

func showPreHandle(priority *int, filterByPriority *bool) func(cmd *cobra.Command, args []string) error {
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
