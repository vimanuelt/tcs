package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"tcs/fileops"
	"tcs/shared"
	"tcs/snapshot"
	"testing"
)

func TestSnapshotAndDeltaIntegration(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "integration_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create initial files
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.bin")
	ioutil.WriteFile(file1, []byte("Initial text content."), 0644)
	ioutil.WriteFile(file2, []byte{0x01, 0x02, 0x03}, 0644)

	// Create initial snapshot
	snapshotFile := filepath.Join(tempDir, "snapshot.json")
	currentFiles, err := fileops.ProcessDirectory(tempDir)
	if err := snapshot.SaveSnapshot(snapshotFile, currentFiles); err != nil {
		t.Fatalf("Failed to save snapshot: %v", err)
	}

	// Modify directory: add and remove files
	file3 := filepath.Join(tempDir, "file3.txt")
	ioutil.WriteFile(file3, []byte("New file content."), 0644)
	os.Remove(file2)

	// Create updated snapshot
	updatedFiles, err := fileops.ProcessDirectory(tempDir)
	updatedSnapshotFile := filepath.Join(tempDir, "updated_snapshot.json")
	if err := snapshot.SaveSnapshot(updatedSnapshotFile, updatedFiles); err != nil {
		t.Fatalf("Failed to save updated snapshot: %v", err)
	}

	// Load previous snapshot
	previousFiles, err := snapshot.LoadSnapshot(snapshotFile)
	if err != nil {
		t.Fatalf("Failed to load previous snapshot: %v", err)
	}

	// Compare snapshots
	added, removed, modified := snapshot.CompareSnapshots(updatedFiles, previousFiles)

	// Filter out snapshot.json from added files
	filteredAdded := []shared.FileInfo{}
	for _, file := range added {
		if filepath.Base(file.Path) != "snapshot.json" {
			filteredAdded = append(filteredAdded, file)
		}
	}

	// Filter out the root directory from modified files
	filteredModified := []shared.FileInfo{}
	for _, file := range modified {
		if file.Path != tempDir {
			filteredModified = append(filteredModified, file)
		}
	}

	// Validate results
	if len(filteredAdded) != 1 || filepath.Base(filteredAdded[0].Path) != "file3.txt" {
		t.Errorf("Expected added file file3.txt, got %v", filteredAdded)
	}

	if len(removed) != 1 || filepath.Base(removed[0].Path) != "file2.bin" {
		t.Errorf("Expected removed file file2.bin, got %v", removed)
	}

	if len(filteredModified) != 0 {
		t.Errorf("Expected no modified files, got %v", filteredModified)
	}
}
