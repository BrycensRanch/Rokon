package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func renameLogFileIfExists(logDir, logFilePath string) {
	if !fileExists(logFilePath) {
		log.Println("No log file to process")
		return // No log file to process
	}

	today := time.Now().Format("2006-01-02") // Date format: YYYY-MM-DD
	backupLogPath := filepath.Join(logDir, fmt.Sprintf("main-%s.log", today))

	if fileExists(backupLogPath) {
		handleExistingBackupLog(backupLogPath, logFilePath)
	} else {
		renameLogFile(logFilePath, backupLogPath)
	}
}

// Handle the situation when the rotated log for this day already exists
func handleExistingBackupLog(backupLogPath, logFilePath string) {
	backupLogFileBytes, err := os.ReadFile(backupLogPath)
	if err != nil {
		log.Printf("Couldn't read %s. Not attempting to append it with latest.log from today.", backupLogPath)
		return
	}

	logFileBytes, err := os.ReadFile(logFilePath)
	if err != nil {
		log.Printf("Couldn't read %s. Skipping the log appending.", logFilePath)
		return
	}

	// Append the contents of the latest log to the rotated log
	combinedLogFileBytes := append(backupLogFileBytes, logFileBytes...)
	err = os.WriteFile(backupLogPath, combinedLogFileBytes, 0o755)
	if err != nil {
		log.Printf("Error writing to %s: %v", backupLogPath, err)
		return
	}

	log.Printf("Appended previous log data to existing log file: %s\n", backupLogPath)
}

// Rename the log file to the rotated name
func renameLogFile(logFilePath, backupLogPath string) {
	err := os.Rename(logFilePath, backupLogPath)
	if err != nil {
		log.Printf("Error renaming log file: %v\n", err)
		return
	}

	log.Printf("Renamed **OLD** latest.log to %s\n", backupLogPath)
}

func debug() {
	debugLoggingEnabled := viper.GetBool("debug")
	if !debugLoggingEnabled {
		return
	}
	log.SetPrefix("Debug")
	// Set the log prefix to indicate it's debug output
	log.SetPrefix("DEBUG: ")

	// Optionally, set the log flags to display timestamp, file, and line number for debugging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("test")
}
