package utils

import "time"

func ResolveDate(value string) (time.Time, error) {
	return time.Now(), nil
}

func SameDate(a, b time.Time) bool {
	a = a.Local()
	b = b.Local()
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}
