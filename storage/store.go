package storage

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

type Record struct {
	Note
}

type Store struct {
	path string
}

func NewStore(path string) *Store {
	return &Store{path: path}
}

func DefaultStore() *Store {
	return NewStore(defaultDataPath())
}

func (s *Store) Path() string {
	return s.path
}

func (s *Store) Append(record Record) error {
	// 当前实现走“读全量再写全量”，逻辑简单，适合这个单文件小项目。
	records, err := s.List()
	if err != nil {
		return err
	}
	records = append(records, record)
	return s.Replace(records)
}

func (s *Store) List() ([]Record, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// 数据文件还没创建时，直接视为空列表。
			return []Record{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return []Record{}, nil
	}

	var records []Record
	// 整个文件就是一个 JSON 数组，直接反序列化即可。
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}
	return records, nil
}

func (s *Store) Replace(records []Record) error {
	// 写文件前先确保 data 目录存在。
	if err := os.MkdirAll(filepathDir(s.path), 0o755); err != nil {
		return err
	}

	// 用缩进格式保存，方便直接打开文件排查问题。
	data, err := json.MarshalIndent(slices.Clone(records), "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')

	return os.WriteFile(s.path, data, 0o644)
}

func defaultDataPath() string {
	return dataPathForDate(nil)
}
