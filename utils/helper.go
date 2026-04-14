package utils

import (
	"dodolist/storage"
	"fmt"
	"strings"
	"time"
)

// 这里定义列表展示时使用的视图结构。
type ViewItem struct {
	Index int
	storage.Record
}

func FormatStatus(completed bool) string {
	// 把布尔完成状态转换成更直观的文本。
	if completed {
		return "done"
	}
	return "todo"
}

func JoinContent(args []string) string {
	// 把命令行剩余参数重新拼成完整待办内容。
	return strings.TrimSpace(strings.Join(args, " "))
}

func ParseIndex(value string) (int, error) {
	// 先复用正整数解析逻辑拿到用户输入索引。
	index, err := ParsePositiveIndex(value)
	if err != nil {
		return 0, err
	}

	// 命令行里的索引从 1 开始，内部切片下标从 0 开始。
	return index - 1, nil
}

func CurrentStore() *storage.Store {
	// 对命令层统一返回默认存储对象。
	return storage.DefaultStore()
}

func BuildViewItems(records []storage.Record) []ViewItem {
	// 先准备好与原始记录数相同容量的展示切片。
	items := make([]ViewItem, 0, len(records))
	for i, record := range records {
		// 提前把展示索引算好，列表排序后也能保留原始编号。
		items = append(items, ViewItem{
			Index:  i + 1,
			Record: record,
		})
	}
	return items
}

func FormatTime(t time.Time) string {
	// 零值时间直接显示占位符，避免界面里出现奇怪的默认时间。
	if t.IsZero() {
		// 零值时间统一显示占位符，避免输出难懂的默认时间。
		return "-"
	}

	// 正常时间统一格式化成固定的本地时间字符串。
	return t.Local().Format("2006-01-02 15:04:05")
}

func ValidatePriority(priority int) error {
	// 当前项目约定优先级不能小于零。
	if priority < 0 {
		return fmt.Errorf("priority must be greater than or equal to 0")
	}
	return nil
}
