// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/koron/go-ssdp"
)

var (
	version                = "0.0.0-SNAPSHOT"
	isPackaged             = "false"
	packageFormat          = "native"
	telemetryOnByDefault   = "true"
	commit                 = "unknown"
	branch                 = "unknown"
	date                   = "unknown"
	logFilePath            = filepath.Join(xdg.DataHome, "rokon", "logs", "latest.log")
	portableDataFolderName = "data"
	tempDirPath            = filepath.Join(xdg.CacheHome, "rokon")
	configDirectoryPath    = filepath.Join(xdg.ConfigHome, "rokon")
)

func main() {
	execDir, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return
	}
	execDir = filepath.Dir(execDir) // Get the directory where the executable is located

	if fileExists(filepath.Join(execDir, "portable.txt")) {
		logFilePath = filepath.Join(execDir, portableDataFolderName, "logs", "latest.log")
		tempDirPath = filepath.Join(execDir, portableDataFolderName, "cache")
		configDirectoryPath = filepath.Join(execDir, portableDataFolderName, "config")
	}
	logDir := filepath.Dir(logFilePath)

	err = os.MkdirAll(filepath.Dir(logFilePath), 0o755)
	if err != nil {
		log.Println("Error creating directory:", err)
		return
	}
	renameLogFileIfExists(logDir, logFilePath)

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
	if err != nil {
		log.Println("Error opening rokon log file:", err)
		return
	}
	defer logFile.Close()
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	ssdpLogger := log.New(logFile, "SSDP: ", log.Ldate|log.Ltime|log.Lshortfile)
	telemetryLogger := log.New(logFile, "TELEMETRY: ", log.Ldate|log.Ltime|log.Lshortfile)
	ssdp.Logger = ssdpLogger

	log.Printf("Config directory: %s", configDirectoryPath)
	log.Printf("Cache directory: %s", tempDirPath)
	log.Printf("Log file: %s", logFilePath)
	log.Println("Starting Rokon. Now with more telemetry!")
	switch runtime.GOOS {
	case "windows", "darwin":
		log.Println("Running on Windows or macOS.")
		// Use GLib to set the GTK_CSD environment variable for Client-Side Decorations
		glib.Setenv("GTK_CSD", "0", true)
		os.Setenv("GTK_CSD", "0")
	default:
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(configDirectoryPath)
	viper.AutomaticEnv()
	for _, dir := range xdg.ConfigDirs {
		viper.AddConfigPath(dir)
	}
	viper.SetEnvPrefix("rokon") // will be uppercased automatically
	viper.SetDefault("telemetry", (telemetryOnByDefault == "1" || telemetryOnByDefault == "true"))
	viper.SetDefault("scanOnStartup", true)

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfigAs(filepath.Join(configDirectoryPath, "config.toml")) // Deploy the default config!
		} else {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	app := gtk.NewApplication("io.github.brycensranch.Rokon", gio.ApplicationHandlesCommandLine)
	app.SetDefault()
	telemetry := viper.GetBool("telemetry")
	if version != "" {
		app.SetVersion(version)
	}
	if telemetry {
		doTelemetry(*telemetryLogger, app, execDir)
	}
	log.Printf("Version %s commit %s branch %s (built on %s)\n", version, commit, branch, date)
	app.ConnectActivate(func() { activate(app) })
	app.ConnectCommandLine(func(commandLine *gio.ApplicationCommandLine) int {
		return activateCommandLine(app, commandLine)
	})
	// Get the command-line arguments
	args := os.Args

	// Check for the flags that should trigger empty args
	// if containsFlag(args, "--version") || containsFlag(args, "--gpplication-service") || containsFlag(args, "--help-gapplication") {
	// 	// Create an empty argument list to pass to app.Run()
	// 	args = []string{}
	// }

	if code := app.Run(args); code > 0 {
		os.Exit(code)
	}
}
