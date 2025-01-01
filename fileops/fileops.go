package fileops

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"sync"
	"tcs/shared"
)

// ProcessDirectory processes files in a directory concurrently.
// It traverses the directory structure, calculates file hashes, and collects metadata.
func ProcessDirectory(dir string) ([]shared.FileInfo, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var files []shared.FileInfo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(path string, info os.FileInfo) {
			defer wg.Done()

			fileInfo := shared.FileInfo{
				Path:        path,
				Type:        "file",
				Size:        info.Size(),
				LastUpdated: info.ModTime().String(),
			}

			if info.IsDir() {
				fileInfo.Type = "directory"
				mu.Lock()
				files = append(files, fileInfo)
				mu.Unlock()
				return
			}

			if hash, err := calculateFileHash(path); err == nil {
				fileInfo.Hash = hash
			}

			mu.Lock()
			files = append(files, fileInfo)
			mu.Unlock()
		}(path, info)

		return nil
	})

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return files, nil
}

// calculateFileHash calculates the SHA-256 hash of a file.
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			hasher.Write(buffer[:n])
		}
		if err != nil {
			break
		}
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
