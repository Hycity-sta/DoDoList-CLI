package storage

import "time"

func ParseDate(value string) (time.Time, error) {
	// 这里统一按固定日期格式解析字符串。
	return time.Parse("2006-01-02", value)
}
