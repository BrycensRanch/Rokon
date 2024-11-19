// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"github.com/koron/go-ssdp"
)

var (
	version              = "0.0.0-SNAPSHOT"
	isPackaged           = "false"
	packageFormat        = "native"
	telemetryOnByDefault = "true"
	commit               = "unknown"
	branch               = "unknown"
	date                 = "unknown"
	logFilePath          = filepath.Join(xdg.DataHome, "rokon", "logs", "latest.log")
	tempDirPath          = filepath.Join(xdg.CacheHome, "rokon")
	configDirectoryPath  = filepath.Join(xdg.ConfigHome, "rokon")
)

func main() {
	execDir, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return
	}
	execDir = filepath.Dir(execDir) // Get the directory where the executable is located

	if fileExists(filepath.Join(execDir, "portable.txt")) {
		logFilePath = filepath.Join(execDir, "data", "logs", "latest.log")
		tempDirPath = filepath.Join(execDir, "data", "cache")
		configDirectoryPath = filepath.Join(execDir, "data", "config")
	}
	logDir := filepath.Dir(logFilePath)

	err = os.MkdirAll(filepath.Dir(logFilePath), 0o755)
	if err != nil {
		log.Println("Error creating directory:", err)
		return
	}

	// Check if the latest.log file exists, and if so, rename it to the date-based file name
	if fileExists(logFilePath) {
		// Generate the new log filename based on the current date
		today := time.Now().Format("2006-01-02") // Date format: YYYY-MM-DD
		backupLogPath := filepath.Join(logDir, fmt.Sprintf("main-%s.log", today))

		if fileExists(backupLogPath) {
			backupLogFileBytes, err := os.ReadFile(backupLogPath)
			// Irresponsibly ignoring errors! Don't do this at home, kids.
			logFileBytes, _ := os.ReadFile(logFilePath)
			combinedLogFileBytes := append(backupLogFileBytes, logFileBytes...)
			if err != nil {
				log.Printf("Couldn't read %s. Not attempting to append it with latest.log from today.", backupLogPath)
			} else {
				os.WriteFile(backupLogPath, combinedLogFileBytes, 0o755)
			}
		} else {
			// Rename the current latest.log to the new file
			err := os.Rename(logFilePath, backupLogPath)
			if err != nil {
				log.Printf("Error renaming log file: %v\n", err)
				return
			}
		}
		log.Printf("Renamed **OLD** latest.log to %s\n", backupLogPath)
	}
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
	if err != nil {
		log.Println("Error opening rokon log file:", err)
		return
	}
	defer logFile.Close()
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set the logger to output to the file
	log.SetOutput(multiWriter)

	// You can also customize the logger with a prefix and timestamp
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// Create a new logger that writes to the file.
	customLogger := log.New(logFile, "SSDP: ", log.Ldate|log.Ltime|log.Lshortfile)
	telemetryLogger := log.New(logFile, "TELEMETRY: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Assign the custom logger to the SSDP.Logger field.
	ssdp.Logger = customLogger
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
			// Handle errors reading the config file
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	app := gtk.NewApplication("io.github.brycensranch.Rokon", gio.ApplicationFlagsNone)
	telemetry := viper.GetBool("telemetry")
	if version != "" {
		app.SetVersion(version)
	}
	if telemetry {
		doTelemetry(telemetryLogger, app, execDir)
	}
	log.Printf("Version %s commit %s branch %s (built on %s)\n", version, commit, branch, date)
	app.ConnectActivate(func() { activate(app) })
	app.ConnectCommandLine(func(commandLine *gio.ApplicationCommandLine) int {
		return activateCommandLine(app, commandLine)
	})
	// Get the command-line arguments
	args := os.Args

	// Check for the flags that should trigger empty args
	if containsFlag(args, "--version") || containsFlag(args, "--gpplication-service") || containsFlag(args, "--help-gapplication") {
		// Create an empty argument list to pass to app.Run()
		args = []string{}
	}

	if code := app.Run(args); code > 0 {
		os.Exit(code)
	}
}

func activateCommandLine(_ *gtk.Application, commandLine *gio.ApplicationCommandLine) int {
	args := commandLine.Arguments() // Get the command-line arguments
	// Check if --version flag is present
	for _, arg := range args {
		if arg == "version" || arg == "--version" {
			// Print version info
			log.Println(applicationInfo())
			return 0 // Return 0 to indicate success
		}
	}
	commandLine.PrintLiteral("HI FROM COMMAND LINE RAHH")
	return 0
}

func activate(app *gtk.Application) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error fetching network interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		// Get the interface status
		status := "down"
		if iface.Flags&net.FlagUp != 0 {
			status = "up"
		}

		// Determine the type of the interface
		var ifaceType string
		switch {
		case iface.Flags&net.FlagLoopback != 0:
			ifaceType = "loopback"
		case strings.Contains(iface.Name, "en") || strings.Contains(iface.Name, "eth"):
			ifaceType = "Ethernet"
		case strings.Contains(iface.Name, "wl"):
			ifaceType = "Wi-Fi"
		default:
			ifaceType = "Unknown"
		}

		// Print interface details
		log.Printf("Interface: %s, Status: %s, Type: %s\n", iface.Name, status, ifaceType)
	}

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Rokon: Control your Roku from your desktop")
	window.SetChild(&gtk.NewLabel("Searching for Rokus on your network...").Widget)
	// Create the main menu
	menu := createMenu(window, app)
	app.SetMenubar(menu)
	window.SetShowMenubar(true)
	window.SetTitle("Rokon: Control your Roku from your desktop")
	window.SetChild(&gtk.NewLabel("Searching for Rokus on your network...").Widget)
	windowWidth := 800
	windowHeight := 400
	window.SetDefaultSize(windowWidth, windowHeight)
	window.SetVisible(true)

	// Event controller setup
	keyController := gtk.NewEventControllerKey()
	keyController.SetName("keyController")
	window.AddController(keyController)

	keyController.Connect("key-pressed", func(controller *gtk.EventControllerKey, code uint) {
		log.Println(controller.Name() + " " + strconv.FormatUint(uint64(code), 10))
		const (
			RightClickCode = uint(93) // Code representing a right-click
		)
		if code == RightClickCode {
			log.Println("right clicked")
		}
	})

	focusController := gtk.NewEventControllerFocus()
	focusController.SetName("focusController")

	focusController.Connect("enter", func() {
		log.Println("Keyboard focus entered!")
	})
	window.AddController(focusController)

	gestureClick := gtk.NewGestureClick()
	gestureClick.SetName("gestureClick")
	gestureClick.Connect("pressed", func(_, numberOfPresses uint) {
		log.Printf("Number of presses %v", numberOfPresses)
	})
	window.AddController(gestureClick)

	// window.Maximize()
	scanOnStartup := viper.GetBool("scanOnStartup")
	// Start searching for Rokus when the app is activated
	var rokuChan chan []ssdp.Service
	if scanOnStartup {
		rokuChan = searchForRokus()
	} else {
		// Empty!
		rokuChan = make(chan []ssdp.Service)
		log.Printf("scanOnStartup is FALSE. Not scanning for Rokus.")
		glib.IdleAdd(func() {
			window.SetChild(&gtk.NewLabel("Welcome to Rokon, to get started, enter your Roku's IP address.\nTo get it's IP address, go into Settings -> Network").Widget)
		})
	}

	// Goroutine that waits for Roku discovery to finish
	go func() {
		discoveredRokus := <-rokuChan // Receive the result from the Roku discovery

		// Use glib.IdleAdd to ensure UI updates happen on the main thread
		glib.IdleAdd(func() {
			if discoveredRokus != nil {
				log.Println("Discovered Rokus:", discoveredRokus)
				window.SetChild(&gtk.NewLabel("Discovered Rokus:").Widget)
				log.Println("Number of goroutines:", runtime.NumGoroutine())

			} else {
				window.SetChild(&gtk.NewLabel("No Rokus discovered via SSDP!").Widget)
			}
		})

		var root Root

		if discoveredRokus != nil {
			client := resty.New()
			resp, err := client.R().
				SetResult(&root).
				Get(discoveredRokus[0].Location + "/")

			if err != nil {
				log.Println("Error:", err)
			} else {
				log.Println("Body:", resp)
			}

			notification := gio.NewNotification("Roku discovered")
			var rokuList []string
			const (
				MaxValue = int(2)
			)
			for i, roku := range discoveredRokus {
				if i < MaxValue {
					rokuList = append(rokuList, fmt.Sprintf("Roku Device %d: %v", i+1, roku.Location))
				}
			}
			if len(discoveredRokus) > MaxValue {
				rokuList = append(rokuList, fmt.Sprintf("...and %d more devices", len(discoveredRokus)-MaxValue))
			}
			rokuListString := strings.Join(rokuList, "\n")
			notification.SetBody(rokuListString)

			url := discoveredRokus[0].Location + "/device-image.png"
			deviceImagePath, err := fetchImageAndDownload(url)
			if err != nil {
				log.Println("Error getting image from URL:", err)
				return
			}
			notification.SetIcon(gio.NewFileIcon(gio.NewFileForPath(deviceImagePath)))
			notification.SetDefaultAction("app.connect-roku")
			notification.SetCategory("device")
			app.SendNotification("roku-discovered", notification)

			// UI update for discovered Rokus
			glib.IdleAdd(func() {
				if discoveredRokus != nil {
					log.Println("Discovered Rokus:", discoveredRokus)
					const spacing = int(5)
					vbox := gtk.NewBox(gtk.OrientationVertical, spacing)
					grid := gtk.NewGrid()
					grid.Attach(&vbox.Widget, 1, 0, 1, 1)
					// vbox.SetMarginTop(10)          // Optional: Add some margin to the top
					// vbox.SetMarginEnd(10)          // Optional: Add some margin to the right
					// vbox.SetHAlign(gtk.AlignEnd)   // Align horizontally to the right (end)
					// vbox.SetVAlign(gtk.AlignStart) // Align vertically to the top (start)
					window.SetChild(grid)

					labelText := fmt.Sprintf("Friendly Name: %s\nIP Address: %s",
						root.Device.FriendlyName, discoveredRokus[0].Location)

					label := gtk.NewLabel(labelText)
					vbox.Append(label) // Add label to the vertical box
					log.Println("Number of goroutines:", runtime.NumGoroutine())

				} else {
					window.SetChild(&gtk.NewLabel("No Rokus discovered via SSDP!").Widget)
				}
			})
		}
	}()
}
