package logging

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSetupLog(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testlog.log")
	if err != nil {
		t.Fatalf("Failed to create temporary log file: %v", err)
	}
	logFileName := tempFile.Name()
	tempFile.Close()
	defer os.Remove(logFileName)

	logFile, err := SetupLog(logFileName)
	if err != nil {
		t.Fatalf("SetupLog failed: %v", err)
	}
	defer logFile.Close()

	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		t.Errorf("Log file %s was not created", logFileName)
	}
}

func TestLogInfo(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testlog.log")
	if err != nil {
		t.Fatalf("Failed to create temporary log file: %v", err)
	}
	logFileName := tempFile.Name()
	tempFile.Close()
	defer os.Remove(logFileName)

	logFile, err := SetupLog(logFileName)
	if err != nil {
		t.Fatalf("SetupLog failed: %v", err)
	}
	defer logFile.Close()

	LogInfo("Test info message")

	content, err := ioutil.ReadFile(logFileName)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if string(content) == "" {
		t.Errorf("Expected log content, but log file is empty")
	}
}
