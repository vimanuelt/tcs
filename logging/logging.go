package logging

import (
	"log"
	"os"
)

// SetupLog initializes the logging system.
// It creates or opens the specified log file and sets it as the output for log messages.
func SetupLog(logFileName string) (*os.File, error) {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logFile, nil
}

// LogInfo logs an informational message.
func LogInfo(format string, v ...interface{}) {
	log.Printf("INFO: "+format, v...)
}

// LogError logs an error message.
func LogError(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}
