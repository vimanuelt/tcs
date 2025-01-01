package snapshot

import (
	"encoding/json"
	"os"
	"tcs/shared"
)

// SaveSnapshot writes the snapshot of the directory structure and file metadata to a file.
func SaveSnapshot(outputFile string, files []shared.FileInfo) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(files)
}

// LoadSnapshot reads a snapshot file and returns its contents as a slice of FileInfo.
func LoadSnapshot(snapshotFile string) ([]shared.FileInfo, error) {
	file, err := os.Open(snapshotFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var files []shared.FileInfo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&files); err != nil {
		return nil, err
	}

	return files, nil
}

// CompareSnapshots compares two snapshots and returns slices of added, removed, and modified files.
func CompareSnapshots(current, previous []shared.FileInfo) (added, removed, modified []shared.FileInfo) {
	previousMap := make(map[string]shared.FileInfo)
	for _, file := range previous {
		previousMap[file.Path] = file
	}

	for _, file := range current {
		if prevFile, exists := previousMap[file.Path]; exists {
			if file.Size != prevFile.Size || file.LastUpdated != prevFile.LastUpdated || file.Hash != prevFile.Hash {
				modified = append(modified, file)
			}
			delete(previousMap, file.Path)
		} else {
			added = append(added, file)
		}
	}

	for _, file := range previousMap {
		removed = append(removed, file)
	}

	return added, removed, modified
}

// SaveDeltaReport writes added, removed, and modified files to a JSON file.
func SaveDeltaReport(deltaFile string, added, removed, modified []shared.FileInfo) error {
	file, err := os.Create(deltaFile)
	if err != nil {
		return err
	}
	defer file.Close()

	report := map[string]interface{}{
		"added":    added,
		"removed":  removed,
		"modified": modified,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(report)
}
