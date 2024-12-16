// SPDX-License-Identifier: AGPL-3.0-or-later
package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"

	"github.com/koron/go-ssdp"
)

var (
	Version                = "0.0.0-SNAPSHOT"
	IsPackaged             = "false"
	PackageFormat          = "native"
	TelemetryOnByDefault   = "true"
	Commit                 = "unknown"
	Branch                 = "unknown"
	Date                   = "unknown"
	LogFilePath            = filepath.Join(xdg.DataHome, "rokon", "logs", "latest.log")
	PortableDataFolderName = "data"
	TempDirPath            = filepath.Join(xdg.CacheHome, "rokon")
	ConfigDirectoryPath    = filepath.Join(xdg.ConfigHome, "rokon")
	ExecDir 			   = ""
	LogFile                *os.File
	Telemetry 			   = (TelemetryOnByDefault == "1" || TelemetryOnByDefault == "true")
)

func Main() {
	ExecDir, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return
	}
	ExecDir = filepath.Dir(ExecDir) // Get the directory where the executable is located

	if FileExists(filepath.Join(ExecDir, "portable.txt")) {
		LogFilePath = filepath.Join(ExecDir, PortableDataFolderName, "logs", "latest.log")
		TempDirPath = filepath.Join(ExecDir, PortableDataFolderName, "cache")
		ConfigDirectoryPath = filepath.Join(ExecDir, PortableDataFolderName, "config")
	}
	logDir := filepath.Dir(LogFilePath)

	err = os.MkdirAll(filepath.Dir(LogFilePath), 0o755)
	if err != nil {
		log.Println("Error creating directory:", err)
		return
	}
	RenameLogFileIfExists(logDir, LogFilePath)

	LogFile, err := os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
	if err != nil {
		log.Println("Error opening rokon log file:", err)
		return
	}
	defer LogFile.Close()
	multiWriter := io.MultiWriter(os.Stdout, LogFile)

	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	ssdpLogger := log.New(LogFile, "SSDP: ", log.Ldate|log.Ltime|log.Lshortfile)
	ssdp.Logger = ssdpLogger

	log.Printf("Config directory: %s", ConfigDirectoryPath)
	log.Printf("Cache directory: %s", TempDirPath)
	log.Printf("Operating System: %s (%s,%s)\n", GetOSRelease(), runtime.GOOS, runtime.GOARCH)
	log.Printf("Log file: %s", LogFilePath)
	log.Println("Starting Rokon. Now with more telemetry!")
	telemetryLogger := log.New(LogFile, "TELEMETRY: ", log.Ldate|log.Ltime|log.Lshortfile)
	Telemetry = viper.GetBool("telemetry")
	if Telemetry {
		DoTelemetry(*telemetryLogger, ExecDir)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(ConfigDirectoryPath)
	viper.AutomaticEnv()
	for _, dir := range xdg.ConfigDirs {
		viper.AddConfigPath(dir)
	}
	viper.SetEnvPrefix("rokon") // will be uppercased automatically
	viper.SetDefault("telemetry", Telemetry)
	viper.SetDefault("scanOnStartup", true)

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfigAs(filepath.Join(ConfigDirectoryPath, "config.toml")) // Deploy the default config!
		} else {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}
	log.Printf("Version %s commit %s branch %s (built on %s)\n", Version, Commit, Branch, Date)
}
