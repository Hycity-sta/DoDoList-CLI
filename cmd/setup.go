package cmd

import (
	"dodolist/i18n"
	"dodolist/storage"
	"dodolist/utils"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// 用于设置根命令
func Setup() *cobra.Command {
	// 在命令包内统一定义根命令和所有子命令
	root := &cobra.Command{
		Use:     "dodolist",
		Short:   i18n.T(i18n.AppShort),
		Long:    i18n.T(i18n.AppLong),
		Args:    cobra.ArbitraryArgs,
		Example: "  dodolist\n  dodolist Buy milk\n  dodolist ok 1\n  dodolist clear\n  dodolist delete 1",
		RunE:    setupHandle,
	}

	// 关掉自动生成的 completion 命令，只保留项目自己定义的命令
	root.CompletionOptions.DisableDefaultCmd = true
	root.SetHelpCommand(Help())

	// 将所有命令挂载到根命令上
	root.AddCommand(Version())
	root.AddCommand(Lang())
	root.AddCommand(Ok())
	root.AddCommand(Clear())
	root.AddCommand(Delete())
	return root
}

func setupHandle(cmd *cobra.Command, args []string) error {
	// 没有内容参数时默认展示待办列表
	if len(args) == 0 {
		return showTodos(cmd, args)
	}
	// 传入内容时直接走新增待办流程
	return createTodo(cmd, args)
}

func createTodo(cmd *cobra.Command, args []string) error {
	// 组合用户输入，构造要写入存储的新待办记录
	store := utils.CurrentStore()
	record := storage.Record{
		Note: storage.Note{
			Content:   utils.JoinContent(args),
			CreatedAt: time.Now(),
		},
	}

	// 追加到当前存储后，再向终端输出创建结果
	if err := store.Append(record); err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", i18n.T(i18n.OutputCreatedTodo, record.Content))
	return nil
}

func showTodos(cmd *cobra.Command, args []string) error {
	// 先从存储里读取全部待办，再按创建时间展示
	store := utils.CurrentStore()
	records, err := store.List()
	if err != nil {
		return err
	}

	items := utils.BuildViewItems(records)
	// 先准备表头和数据行，再按显示宽度计算列宽，避免中英文混排错位
	headers := []string{
		i18n.T(i18n.CmdShowHeaderIndex),
		i18n.T(i18n.CmdShowHeaderCreatedAt),
		i18n.T(i18n.CmdShowHeaderStatus),
		i18n.T(i18n.CmdShowHeaderTodo),
	}
	rows := make([][]string, 0, len(items))
	for _, item := range items {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.Index),
			utils.FormatTime(item.CreatedAt),
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

	// 统一用补齐后的列宽输出表格
	writeRow := func(values []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s  %s  %s  %s\n",
			utils.PadRightDisplay(values[0], widths[0]),
			utils.PadRightDisplay(values[1], widths[1]),
			utils.PadRightDisplay(values[2], widths[2]),
			utils.PadRightDisplay(values[3], widths[3]),
		)
	}

	separator := []string{
		strings.Repeat("-", widths[0]),
		strings.Repeat("-", widths[1]),
		strings.Repeat("-", widths[2]),
		strings.Repeat("-", widths[3]),
	}

	writeRow(headers)
	writeRow(separator)
	for _, row := range rows {
		writeRow(row)
	}
	return nil
}
