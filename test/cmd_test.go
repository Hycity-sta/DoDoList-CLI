package test

import (
	"bytes"
	"dodolist/cmd"
	"dodolist/storage"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestListSortsByPriority(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")
	store := storage.NewStore(path)
	err := store.Replace([]storage.Record{
		{Note: storage.Note{Content: "low", Priority: 1, CreatedAt: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)}},
		{Note: storage.Note{Content: "high", Priority: 3, CreatedAt: time.Date(2026, 4, 15, 11, 0, 0, 0, time.UTC)}},
	})
	if err != nil {
		t.Fatalf("replace failed: %v", err)
	}

	root := cmd.RootForTest(path)
	output := &bytes.Buffer{}
	root.SetOut(output)
	root.SetErr(output)
	root.SetArgs([]string{"list", "--sort"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	text := output.String()
	if strings.Index(text, "high") > strings.Index(text, "low") {
		t.Fatalf("expected high priority todo first, got output:\n%s", text)
	}
}
