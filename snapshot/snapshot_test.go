package snapshot

import (
	"io/ioutil"
	"os"
	"tcs/shared"
	"testing"
)

func TestSaveSnapshot(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "snapshot.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testFiles := []shared.FileInfo{
		{Path: "file1.txt", Type: "text", Size: 123, LastUpdated: "2025-01-01T00:00:00Z"},
		{Path: "file2.bin", Type: "binary", Size: 456, LastUpdated: "2025-01-01T00:00:00Z"},
	}

	if err := SaveSnapshot(tempFile.Name(), testFiles); err != nil {
		t.Fatalf("SaveSnapshot failed: %v", err)
	}

	content, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read snapshot file: %v", err)
	}

	if len(content) == 0 {
		t.Fatalf("Snapshot file is empty")
	}
}
