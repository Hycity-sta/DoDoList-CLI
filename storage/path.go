package storage

import (
	"os"
	"path/filepath"
)

func filepathDir(path string) string {
	// 统一封装目录提取，后面如果要替换实现更方便。
	return filepath.Dir(path)
}

func filepathJoin(elem ...string) string {
	// 统一封装路径拼接，避免各处直接依赖标准库调用细节。
	return filepath.Join(elem...)
}

func filepathClean(path string) string {
	// 统一清理路径格式，避免出现多余分隔符。
	return filepath.Clean(path)
}

func dataPathForDate(_ interface{}) string {
	// 先尝试拿到当前可执行文件所在目录。
	exe, err := os.Executable()
	if err != nil {
		// 如果拿不到可执行文件路径，就退回到相对路径。
		return filepathJoin("data", "todos.json")
	}

	// 正常情况下把数据文件放在程序目录下的 data 目录里。
	return filepathJoin(filepathDir(exe), "data", "todos.json")
}
