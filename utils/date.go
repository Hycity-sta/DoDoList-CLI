package utils

import "time"

func ResolveDate(value string) (time.Time, error) {
	// 这里暂时直接返回当前时间，后面再按业务扩展相对日期解析。
	return time.Now(), nil
}

func SameDate(a, b time.Time) bool {
	// 先统一转成本地时间，避免时区差异影响日期判断。
	a = a.Local()
	b = b.Local()

	// 再按年、月、日逐项比较是否是同一天。
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}
