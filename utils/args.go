package utils

import (
	"fmt"
	"strconv"
)

func ParsePositiveIndex(value string) (int, error) {
	index, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid index %q", value)
	}
	if index <= 0 {
		return 0, fmt.Errorf("index must be greater than 0")
	}
	return index, nil
}
