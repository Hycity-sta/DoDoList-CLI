package storage

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

// 这里把具体待办包一层，方便后面扩展额外元数据。
type Record struct {
	Note
}

// 这里封装单文件存储对象，统一管理读写路径。
type Store struct {
	path string
}

func NewStore(path string) *Store {
	// 用显式路径创建一个新的存储实例。
	return &Store{path: path}
}

func DefaultStore() *Store {
	// 默认存储直接落到程序目录下的数据文件。
	return NewStore(defaultDataPath())
}

func (s *Store) Path() string {
	// 对外暴露当前存储使用的文件路径。
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
	// 先尝试把整份数据文件读进内存。
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// 数据文件还没创建时，直接视为空列表。
			return []Record{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		// 空文件直接按空列表处理。
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

	// 先把待办列表序列化成便于阅读的 JSON 文本。
	// 用缩进格式保存，方便直接打开文件排查问题。
	data, err := json.MarshalIndent(slices.Clone(records), "", "  ")
	if err != nil {
		return err
	}

	// 手动补一个换行，让数据文件在编辑器里看起来更舒服。
	data = append(data, '\n')

	// 最后把完整内容写回磁盘。
	return os.WriteFile(s.path, data, 0o644)
}

func defaultDataPath() string {
	// 默认数据文件路径交给路径工具统一计算。
	return dataPathForDate(nil)
}
