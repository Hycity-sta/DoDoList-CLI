package test

import (
	"dodolist/storage"
	"path/filepath"
	"testing"
	"time"
)

func TestStoreAppendAndList(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")
	store := storage.NewStore(path)

	err := store.Append(storage.Record{
		Note: storage.Note{
			Content:   "write tests",
			Priority:  2,
			CreatedAt: time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC),
		},
	})
	if err != nil {
		t.Fatalf("append failed: %v", err)
	}

	records, err := store.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}
	if records[0].Content != "write tests" {
		t.Fatalf("unexpected content: %q", records[0].Content)
	}
	if records[0].Priority != 2 {
		t.Fatalf("unexpected priority: %d", records[0].Priority)
	}
}
