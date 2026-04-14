package utils

import (
	"dodolist/i18n"
	"fmt"
	"strconv"
)

func ParsePositiveIndex(value string) (int, error) {
	// 先把命令行里的索引文本转换成整数。
	index, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(i18n.T(i18n.ErrInvalidIndex, value))
	}

	// 再检查索引必须大于零，避免出现无效下标。
	if index <= 0 {
		return 0, fmt.Errorf(i18n.T(i18n.ErrIndexPositive))
	}
	return index, nil
}
