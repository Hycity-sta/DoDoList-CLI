package storage

import (
	"os"
	"path/filepath"
)

func filepathDir(path string) string {
	return filepath.Dir(path)
}

func filepathJoin(elem ...string) string {
	return filepath.Join(elem...)
}

func filepathClean(path string) string {
	return filepath.Clean(path)
}

func dataPathForDate(_ interface{}) string {
	exe, err := os.Executable()
	if err != nil {
		return filepathJoin("data", "todos.json")
	}
	return filepathJoin(filepathDir(exe), "data", "todos.json")
}
