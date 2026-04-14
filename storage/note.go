package storage

import "time"

type Note struct {
	// 单条待办的核心字段都落在这里，方便序列化到一个 JSON 文件里。
	Content   string    `json:"content"`
	Priority  int       `json:"priority"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
