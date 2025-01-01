package main

import (
	"fmt"
	"os"

	"tcs/config"
	"tcs/fileops"
	"tcs/logging"
	"tcs/snapshot"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logging
	logFile, err := logging.SetupLog("tcs.log")
	if err != nil {
		fmt.Printf("ERROR: Unable to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logging.LogInfo("Mode: %s", cfg.Mode)
	logging.LogInfo("Directory: %s", cfg.Directory)
	logging.LogInfo("Output File: %s", cfg.OutputFile)
	logging.LogInfo("Previous Snapshot File: %s", cfg.PreviousSnapshotFile)
	logging.LogInfo("Delta File: %s", cfg.DeltaFile)

	// Ensure the directory exists
	if _, err := os.Stat(cfg.Directory); os.IsNotExist(err) {
		logging.LogError("Directory %s does not exist", cfg.Directory)
		os.Exit(1)
	}

	if cfg.Mode == "snapshot" {
		// Process the directory and generate a snapshot
		currentFiles, err := fileops.ProcessDirectory(cfg.Directory)
		if err != nil {
			logging.LogError("Failed to process directory: %v", err)
			os.Exit(1)
		}

		if err := snapshot.SaveSnapshot(cfg.OutputFile, currentFiles); err != nil {
			logging.LogError("Failed to save snapshot: %v", err)
			os.Exit(1)
		}

		logging.LogInfo("Snapshot saved successfully.")
		return
	}

	if cfg.Mode == "delta" {
		// Process the directory and generate a delta report
		currentFiles, err := fileops.ProcessDirectory(cfg.Directory)
		if err != nil {
			logging.LogError("Failed to process directory: %v", err)
			os.Exit(1)
		}

		previousFiles, err := snapshot.LoadSnapshot(cfg.PreviousSnapshotFile)
		if err != nil {
			logging.LogError("Failed to load previous snapshot: %v", err)
			os.Exit(1)
		}

		added, removed, modified := snapshot.CompareSnapshots(currentFiles, previousFiles)

		if err := snapshot.SaveDeltaReport(cfg.DeltaFile, added, removed, modified); err != nil {
			logging.LogError("Failed to save delta report: %v", err)
			os.Exit(1)
		}

		logging.LogInfo("Delta report saved successfully.")
		return
	}

	logging.LogError("Invalid mode: %s. Use 'snapshot' or 'delta'.", cfg.Mode)
	os.Exit(1)
}
