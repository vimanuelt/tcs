package fileops

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestProcessDirectory(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.bin")
	ioutil.WriteFile(file1, []byte("This is a text file."), 0644)
	ioutil.WriteFile(file2, []byte{0x00, 0x01, 0x02}, 0644)

	// Process the directory
	files, err := ProcessDirectory(tempDir)
	if err != nil {
		t.Fatalf("ProcessDirectory failed: %v", err)
	}

	// Filter out the root directory and unexpected files
	expectedFiles := map[string]bool{
		"file1.txt": true,
		"file2.bin": true,
	}
	for _, file := range files {
		// Skip the root directory
		if file.Path == tempDir {
			continue
		}
		if !expectedFiles[filepath.Base(file.Path)] {
			t.Errorf("Unexpected file found: %s", file.Path)
		}
		delete(expectedFiles, filepath.Base(file.Path))
	}

	if len(expectedFiles) != 0 {
		t.Errorf("Expected files missing: %v", expectedFiles)
	}
}
