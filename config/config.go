package config

import (
	"flag"
)

// Config holds the application's runtime configuration.
// It stores user-provided options parsed from command-line arguments.
type Config struct {
	Mode                 string // Mode of operation: snapshot or delta
	Directory            string // Directory to process
	OutputFile           string // Path to save the snapshot file
	PreviousSnapshotFile string // Path to the previous snapshot file for delta mode
	DeltaFile            string // Path to save the delta report file
}

// LoadConfig parses and returns the configuration from command-line flags.
// It defines and processes command-line arguments for the application.
func LoadConfig() *Config {
	mode := flag.String("mode", "snapshot", "Mode of operation: snapshot or delta")
	dir := flag.String("directory", "./", "Directory to process")
	outputFile := flag.String("output-file", "current_file_structure_and_content.json", "Path for the output snapshot file")
	previousSnapshotFile := flag.String("previous-snapshot", "previous_snapshot.json", "Path for the previous snapshot file")
	deltaFile := flag.String("delta-file", "delta_report.json", "Path for the delta report file")

	flag.Parse()

	return &Config{
		Mode:                 *mode,
		Directory:            *dir,
		OutputFile:           *outputFile,
		PreviousSnapshotFile: *previousSnapshotFile,
		DeltaFile:            *deltaFile,
	}
}
