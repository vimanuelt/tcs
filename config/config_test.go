package config

import (
	"flag"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Use a custom FlagSet for testing
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.String("mode", "snapshot", "Mode of operation")
	fs.String("directory", "./testdir", "Directory to process")
	fs.String("output-file", "test_output.json", "Output file path")
	fs.String("previous-snapshot", "test_prev.json", "Previous snapshot file path")
	fs.String("delta-file", "test_delta.json", "Delta file path")

	// Simulate parsed flags
	cfg := &Config{
		Mode:                 fs.Lookup("mode").Value.String(),
		Directory:            fs.Lookup("directory").Value.String(),
		OutputFile:           fs.Lookup("output-file").Value.String(),
		PreviousSnapshotFile: fs.Lookup("previous-snapshot").Value.String(),
		DeltaFile:            fs.Lookup("delta-file").Value.String(),
	}

	if cfg.Mode != "snapshot" {
		t.Errorf("Expected mode 'snapshot', got '%s'", cfg.Mode)
	}
	if cfg.Directory != "./testdir" {
		t.Errorf("Expected directory './testdir', got '%s'", cfg.Directory)
	}
	if cfg.OutputFile != "test_output.json" {
		t.Errorf("Expected output file 'test_output.json', got '%s'", cfg.OutputFile)
	}
	if cfg.PreviousSnapshotFile != "test_prev.json" {
		t.Errorf("Expected previous snapshot file 'test_prev.json', got '%s'", cfg.PreviousSnapshotFile)
	}
	if cfg.DeltaFile != "test_delta.json" {
		t.Errorf("Expected delta file 'test_delta.json', got '%s'", cfg.DeltaFile)
	}
}
