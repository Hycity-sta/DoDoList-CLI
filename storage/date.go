package storage

import "time"

func ParseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}
