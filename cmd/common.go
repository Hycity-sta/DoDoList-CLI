package cmd

import (
	"dodolist/storage"
	"dodolist/utils"
	"fmt"
	"strings"
	"time"
)

type todoView struct {
	Index int
	storage.Record
}

func joinContent(args []string) string {
	return strings.TrimSpace(strings.Join(args, " "))
}

func parseIndexArg(value string) (int, error) {
	index, err := utils.ParsePositiveIndex(value)
	if err != nil {
		return 0, err
	}
	// 命令行里的索引从 1 开始，内部切片下标从 0 开始。
	return index - 1, nil
}

func mustRecords(store *storage.Store) ([]storage.Record, error) {
	return store.List()
}

func currentStore() *storage.Store {
	// 测试环境优先使用覆盖路径，正常运行时走默认数据文件。
	if storePathOverride != "" {
		return storage.NewStore(storePathOverride)
	}
	return storage.DefaultStore()
}

func recordView(records []storage.Record) []todoView {
	items := make([]todoView, 0, len(records))
	for i, record := range records {
		// 提前把展示索引算好，列表排序后也能保留原始编号。
		items = append(items, todoView{
			Index:  i + 1,
			Record: record,
		})
	}
	return items
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		// 零值时间统一显示占位符，避免输出难懂的默认时间。
		return "-"
	}
	return t.Local().Format("2006-01-02 15:04:05")
}

func validatePriority(priority int) error {
	if priority < 0 {
		return fmt.Errorf("priority must be greater than or equal to 0")
	}
	return nil
}
