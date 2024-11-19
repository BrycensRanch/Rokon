// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"golang.org/x/mod/semver"

	"github.com/brycensranch/go-aptabase/pkg/aptabase/v1"
	"github.com/brycensranch/go-aptabase/pkg/osinfo/v1"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"github.com/koron/go-ssdp"
)

var aptabaseClient *aptabase.Client // Package-level variable

// Root represents the root element of the XML.
type Root struct {
	XMLName     xml.Name    `xml:"root"`
	SpecVersion SpecVersion `xml:"specVersion"`
	Device      Device      `xml:"device"`
}

// SpecVersion holds the major and minor version numbers.
type SpecVersion struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

// Device represents the device details.
type Device struct {
	DeviceType       string      `xml:"deviceType"`
	FriendlyName     string      `xml:"friendlyName"`
	Manufacturer     string      `xml:"manufacturer"`
	ManufacturerURL  string      `xml:"manufacturerURL"`
	ModelDescription string      `xml:"modelDescription"`
	ModelName        string      `xml:"modelName"`
	ModelNumber      string      `xml:"modelNumber"`
	ModelURL         string      `xml:"modelURL"`
	SerialNumber     string      `xml:"serialNumber"`
	UDN              string      `xml:"UDN"`
	IconList         IconList    `xml:"iconList"`
	ServiceList      ServiceList `xml:"serviceList"`
}

// IconList holds a list of icons.
type IconList struct {
	Icons []Icon `xml:"icon"`
}

// Icon represents an individual icon.
type Icon struct {
	MimeType string `xml:"mimetype"`
	Width    int    `xml:"width"`
	Height   int    `xml:"height"`
	Depth    int    `xml:"depth"`
	URL      string `xml:"url"`
}

// ServiceList holds a list of services.
type ServiceList struct {
	Services []Service `xml:"service"`
}

// Service represents an individual service.
type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceID   string `xml:"serviceId"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
	SCPDURL     string `xml:"SCPDURL"`
}

// Structure to hold GitHub release information.
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

func getOSRelease() string {
	osName, osVersion := osinfo.GetOSInfo()
	return fmt.Sprintf("%s %s", osName, osVersion)
}

func createEvent(eventName string, eventData map[string]interface{}) {
	event := aptabase.EventData{
		EventName: eventName,
		Props:     eventData,
	}
	aptabaseClient.TrackEvent(event)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

var (
	version              = "0.0.0-SNAPSHOT"
	isPackaged           = "false"
	packageFormat        = "native"
	telemetryOnByDefault = "true"
	commit               = "unknown"
	branch               = "unknown"
	date                 = "unknown"
	logFilePath          = filepath.Join(xdg.DataHome, "rokon", "logs", "latest.log")
	tempDir              = filepath.Join(xdg.CacheHome, "rokon")
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
		tempDir = filepath.Join(execDir, "data", "cache")
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

	// Assign the custom logger to the SSDP.Logger field.
	ssdp.Logger = customLogger
	telemetryLogger := log.New(logFile, "TELEMETRY: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Printf("Log file: %s", logFilePath)
	log.Println("Starting Rokon. Now with more telemetry!")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	configDirectoryPath := filepath.Join(xdg.ConfigHome, "rokon")
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
	telemetry := viper.GetBool("telemetry")
	if telemetry {
	switch runtime.GOOS {
	case "windows", "darwin":
		telemetryLogger.Println("Running on Windows or macOS.")
		// Use GLib to set the GTK_CSD environment variable for Client-Side Decorations
		glib.Setenv("GTK_CSD", "0", true)
		os.Setenv("GTK_CSD", "0")
	default:
	}

	if packageFormat == "detect" {
		_, err = os.Stat("/.dockerenv")

		switch {
		case os.Getenv("APPIMAGE") != "":
			packageFormat = "AppImage"
		case err == nil:
			packageFormat = "docker"
		case os.Getenv("CONTAINER") == "oci":
			packageFormat = "docker"

		case fileExists(filepath.Join(execDir, "share", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "share", "packageFormat"))
			if err == nil {
				packageFormat = string(content)
			}
		case fileExists(filepath.Join(execDir, "usr", "share", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "packageFormat"))
			if err == nil {
				packageFormat = string(content)
			}
		case fileExists(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat"))
			if err == nil {
				packageFormat = string(content)
			}
		case fileExists(filepath.Join(execDir, "usr", "local", "share", "rokon", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat"))
			if err == nil {
				packageFormat = string(content)
			}
		default:
			// Set to "native" if no valid package format is detected
			packageFormat = "native"
		}

		log.Println("Detected package format:", packageFormat)
	}
}

	app := gtk.NewApplication("io.github.brycensranch.Rokon", gio.ApplicationFlagsNone)
	if telemetry {
	aptabaseClient = aptabase.NewClient("A-US-0332858461", version, uint64(133), false, "")
	aptabaseClient.Logger = telemetryLogger
	}
	if version != "" {
		app.SetVersion(version)
	}
	log.Printf("Version %s commit %s branch %s (built on %s)\n", version, commit, branch, date)
	if telemetry {

	switch runtime.GOOS {
	case "linux":
		release := getOSRelease()
		arch := runtime.GOARCH
		desktop := os.Getenv("XDG_CURRENT_DESKTOP")
		sessionType := os.Getenv("XDG_SESSION_TYPE")

		kdeSessionVersion := ""
		if desktop == "KDE" {
			kdeSessionVersion = os.Getenv("KDE_SESSION_VERSION")
		}

		telemetryLogger.Printf("Running on Linux. Specifically: %s %s with %s %s %s and %s\n",
			release, arch, desktop, os.Getenv("DESKTOP_SESSION"), kdeSessionVersion, sessionType)

		createEvent("linux_run", map[string]interface{}{
			"arch":           arch,
			"desktop":        desktop,
			"desktopVersion": kdeSessionVersion,
			"sessionType":    sessionType,
			"packageFormat":  packageFormat,
		})

		container := os.Getenv("container")
		if container != "" && container == "flatpak" {
			telemetryLogger.Println("Running from a Flatpak")
			createEvent("flatpak_run", map[string]interface{}{
				"flatpak":        container,
				"flatpakVersion": version, // Replace with your app version logic
			})
		} else if snap := os.Getenv("SNAP"); snap != "" {
			telemetryLogger.Println("Running from a Snap")
			createEvent("snap_run", map[string]interface{}{
				"snap":        snap,
				"snapVersion": version, // Replace with your app version logic
			})
		} else if appImage := os.Getenv("APPIMAGE"); appImage != "" {
			telemetryLogger.Println("Running from an AppImage")
			firejail := isRunningWithFirejail()

			if firejail {
				telemetryLogger.Println("Running from an AppImage with firejail")
			}

			createEvent("appimage_run", map[string]interface{}{
				"appimage":           filepath.Base(appImage),
				"appimageVersion":    version, // Replace with your app version logic
				"firejail":           firejail,
				"desktopIntegration": os.Getenv("DESKTOPINTEGRATION"),
			})
		} else if isPackaged == "true" {
			telemetryLogger.Println("Running from a native package")
			createEvent("native_run", map[string]interface{}{
				"nativeVersion": version, // Replace with your app version logic
				"path":          path.Base(os.Args[0]),
				"packageFormat": packageFormat,
			})
		}
	case "windows":
		release := getOSRelease()
		arch := runtime.GOARCH
		telemetryLogger.Printf("Running on Windows. Specifically: %s %s\n",
			release, arch)

		if packageFormat == "portable" {
			telemetryLogger.Println("Running from a portable executable")
		}

		createEvent("windows_run", map[string]interface{}{
			"arch":          arch,
			"version":       version, // Replace with your app version logic
			"packageFormat": packageFormat,
		})
	case "darwin":
		release := getOSRelease()
		arch := runtime.GOARCH
		telemetryLogger.Printf("Running on macOS. Specifically: %s %s\n",
			release, arch)

		createEvent("macos_run", map[string]interface{}{
			"arch":          arch,
			"mas":           os.Getenv("MAS"),
			"version":       version, // Replace with your app version logic
			"path":          path.Base(os.Args[0]),
			"packageFormat": packageFormat,
		})
	default:
		telemetryLogger.Printf("Unsupported telemetry platform: %s %s %s. However, the application will continue.\n",
			runtime.GOOS, getOSRelease(), runtime.GOARCH)
		createEvent("unsupported_platform", map[string]interface{}{
			"platform":      runtime.GOOS,
			"arch":          runtime.GOARCH,
			"version":       version, // Replace with your app version logic
			"path":          path.Base(os.Args[0]),
			"packageFormat": packageFormat,
		})
	}
}
	app.ConnectActivate(func() { activate(app) })
	app.ConnectCommandLine(func(commandLine *gio.ApplicationCommandLine) int {
		return activateCommandLine(app, commandLine)
	})
	if telemetry {
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	aptabaseClient.Quit = true
	aptabaseClient.Stop()
	}
	if code := app.Run(os.Args); code > 0 {
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

func applicationInfo() string {
	qualifier := func() string {
		switch {
		case os.Getenv("SNAP") != "":
			return " (Snap)"
		case os.Getenv("container") == "flatpak":
			return " (Flatpak)"
		case os.Getenv("APPIMAGE") != "":
			return " (AppImage)"
		case os.Getenv("CONTAINER") != "":
			return " (Container)"
		case strings.Contains(version, "SNAPSHOT"):
			return " (Development)"
		default:
			return ""
		}
	}()
	return "Rokon" + qualifier
}

func isRunningWithFirejail() bool {
	appImage := os.Getenv("APPIMAGE")
	appDir := os.Getenv("APPDIR")
	return (appImage != "" && contains(appImage, "/run/firejail")) ||
		(appDir != "" && contains(appDir, "/run/firejail"))
}

// Helper function to check if a string contains a substring.
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Search for Rokus asynchronously and return via channel.
func searchForRokus() chan []ssdp.Service {
	resultChan := make(chan []ssdp.Service)

	go func() {
		defer close(resultChan)

		discoveredRokus, err := ssdp.Search("roku:ecp", 1, "")
		if err != nil {
			log.Println("Error discovering Rokus:", err)
			return
		}

		if discoveredRokus != nil {
			resultChan <- discoveredRokus // Send results back to the main thread
			// Deduplicate based on LOCATION
			// Needed because the SSDP code runs on EVERY interface :)
			// So, if you have WiFi and Ethernet enabled, it will show two callbacks from your Roku TV.
			// The code is just that *good*
			locationMap := make(map[string]ssdp.Service)
			for _, roku := range discoveredRokus {
				locationMap[roku.Location] = roku
			}

			// Convert map back to a slice
			uniqueRokus := make([]ssdp.Service, 0, len(locationMap))
			for _, roku := range locationMap {
				uniqueRokus = append(uniqueRokus, roku)
			}
			resultChan <- uniqueRokus // Send results back to the main thread
		} else {
			resultChan <- nil // No Rokus found, send nil
		}
	}()

	return resultChan
}

// Show the "About" window.
func showAboutWindow(mainWindow *gtk.ApplicationWindow, app *gtk.Application) {
	aboutWindow := gtk.NewAboutDialog()
	aboutWindow.SetApplication(app)
	aboutWindow.SetProgramName(applicationInfo())
	aboutWindow.SetVersion(app.Version())
	aboutWindow.SetComments("Control your Roku TV from your desktop")
	aboutWindow.SetWebsite("https://github.com/BrycensRanch/Rokon")
	aboutWindow.SetWebsiteLabel("GitHub")
	//nolint:gosec // In GTK We trust.
	aboutWindow.SetSystemInformation(
		fmt.Sprintf("OS: %s (%s,%s)\n", getOSRelease(), runtime.GOOS, runtime.GOARCH) + fmt.Sprintf("Go: %s\n", runtime.Version()) + fmt.Sprintf("GTK: %d.%d.%d", int(gtk.GetMajorVersion()), int(gtk.GetMinorVersion()), int(gtk.GetMicroVersion())),
	)
	aboutWindow.SetCopyright("©️ 2024 Brycen G and contributors, but mostly Brycen")
	aboutWindow.SetWrapLicense(true)
	aboutWindow.SetModal(false)
	aboutWindow.SetDestroyWithParent(true)

	switch {
	case os.Getenv("SNAP") != "":
		image := gtk.NewImageFromFile(os.Getenv("SNAP") + "/meta/gui/icon.png")
		if image != nil {
			logo := image.Paintable()
			if logo != nil {
				aboutWindow.SetLogo(logo)
			}
		}
	case os.Getenv("FLATPAK") != "":
		image := gtk.NewImageFromFile("/app/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png")
		if image != nil {
			logo := image.Paintable()
			if logo != nil {
				aboutWindow.SetLogo(logo)
			}
		}
	default:
		// Assume native packaging
		aboutWindow.SetLogoIconName("io.github.brycensranch.Rokon")

		if os.Getenv("CONTAINER") != "" {
			log.Println("Running in a container, the logo icon may not be displayed due to wrong path")
		}
	}
	// aboutWindow.SetAuthors([]string{"Brycen G. (BrycensRanch)"})
	aboutWindow.SetLicenseType(gtk.LicenseAGPL30)

	aboutWindow.Present()
	aboutWindow.Focus()
}

// Create the main menu.
func createMenu(window *gtk.ApplicationWindow, app *gtk.Application) *gio.Menu {
	menu := gio.NewMenu()

	// Create "Example" menu item
	exampleMenu := gio.NewMenuItem("Example", "example")
	exampleSubMenu := gio.NewMenu()

	// "About This App" menu item
	aboutMenuItem := gio.NewMenuItem("About This App", "app.about")
	aboutMenuItem.Connect("activate", func() {
		showAboutWindow(window, app)
	})

	// "Check For Updates" menu item
	updateMenuItem := gio.NewMenuItem("Check For Updates", "app.check-for-updates")

	// "Quit" menu item
	quitMenuItem := gio.NewMenuItem("Quit", "quit")
	quitMenuItem.Connect("activate", func() {
		app.Quit()
	})

	aboutAction := gio.NewSimpleAction("about", nil)
	aboutAction.Connect("activate", func() {
		showAboutWindow(window, app)
	})
	app.AddAction(aboutAction)
	exampleSubMenu.AppendItem(aboutMenuItem)

	// Add "Check For Updates" action
	checkForUpdatesAction := gio.NewSimpleAction("check-for-updates", nil)
	checkForUpdatesAction.Connect("activate", func() {
		const url = "https://api.github.com/repos/BrycensRanch/Rokon/releases/latest"

		// Create a Resty client
		client := resty.New()

		// Make the request
		resp, err := client.R().SetResult(&GitHubRelease{}).Get(url)
		if err != nil {
			log.Printf("Update Check Failed. Error fetching release information: %v", err)
			return
		}

		if resp.StatusCode() != 200 {
			log.Printf("Update Check Failed. Unable to fetch release info: " + resp.Status())
			return
		}

		release := resp.Result().(*GitHubRelease)
		latestReleaseVersion := release.TagName
		appVersion := app.Version()
		if semver.Compare(latestReleaseVersion, appVersion) > 0 {
			log.Printf("A new version is available: " + latestReleaseVersion)
		} else if semver.Compare(latestReleaseVersion, appVersion) == 0 {
			log.Printf("Rokon is up to date!")
		} else {
			log.Printf("You're ahead of the latest release", "You are using a newer version than the available release.")
		}
	})

	app.AddAction(checkForUpdatesAction)
	exampleSubMenu.AppendItem(updateMenuItem)

	// Add "Quit" action
	quitAction := gio.NewSimpleAction("quit", nil)
	quitAction.Connect("activate", func() {
		app.Quit()
	})
	app.AddAction(quitAction)
	exampleSubMenu.AppendItem(quitMenuItem)

	exampleMenu.SetSubmenu(exampleSubMenu)
	menu.AppendItem(exampleMenu)

	return menu
}

// Function to show a dialog with the specified title and message.
func showDialog(title, message string, app *gtk.Application) {
	theWindow := gtk.NewWindow()
	dialog := gtk.NewMessageDialog(
		theWindow,
		gtk.DialogDestroyWithParent,
		gtk.MessageInfo,
		gtk.ButtonsNone,
	)

	dialog.SetTitle(title)
	dialog.SetApplication(app)
	dialog.SetModal(true)
	dialog.SetChild(gtk.NewLabel(message))
	theWindow.Show()
	dialog.Show()
}

func fetchImageAsPaintable(url string) (string, error) {
	client := resty.New()
	resp, err := client.SetOutputDirectory(tempDir).EnableTrace().R().
		// SetDebug(true).
		EnableTrace().
		SetOutput(filepath.Join(tempDir, "device-image.png")).
		Get(url)
	if err != nil {
		return "", err
	}
	successfulHTTPCode := 200
	if resp.StatusCode() != successfulHTTPCode {
		return "", fmt.Errorf("failed to get image: status code %d", resp.StatusCode())
	}
	imagePath := filepath.Join(tempDir, "device-image.png")
	log.Println(imagePath)
	// image := gtk.NewImageFromFile(imagePath)

	return imagePath, nil
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
		fmt.Printf("Interface: %s, Status: %s, Type: %s\n", iface.Name, status, ifaceType)
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
		println(controller.Name() + " " + strconv.FormatUint(uint64(code), 10))
		const (
			RightClickCode = uint(93) // Code representing a right-click
		)
		if code == RightClickCode {
			println("right clicked")
		}
	})

	focusController := gtk.NewEventControllerFocus()
	focusController.SetName("focusController")

	focusController.Connect("enter", func() {
		println("Keyboard focus entered!")
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

		// Perform the request and unmarshal directly into the Root struct
		var root Root

		// Once Roku discovery completes, run Resty logic
		if discoveredRokus != nil {
			client := resty.New()
			resp, err := client.R().
				SetResult(&root). // Set the result to automatically unmarshal the response
				Get(discoveredRokus[0].Location + "/")

			if err != nil {
				log.Println("Error:", err)
			} else {
				log.Println("Trace Info:", resp.Request.TraceInfo())
				log.Println("Status Code:", resp.StatusCode())
				log.Println("Status:", resp.Status())
				log.Println("Proto:", resp.Proto())
				log.Println("Time:", resp.Time())
				log.Println("Received At:", resp.ReceivedAt())
				log.Println("Body:", resp)
			}

			notification := gio.NewNotification("Roku discovered")
			var rokuList []string
			const (
				MaxValue = int(3) // Maximum allowed Rokus to display in notification
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
			imagePath, err := fetchImageAsPaintable(url)
			if err != nil {
				log.Println("Error getting image from URL:", err)
				return
			}
			notification.SetIcon(gio.NewFileIcon(gio.NewFileForPath(imagePath)))
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
					vbox.SetMarginTop(10)          // Optional: Add some margin to the top
					vbox.SetMarginEnd(10)          // Optional: Add some margin to the right
					vbox.SetHAlign(gtk.AlignEnd)   // Align horizontally to the right (end)
					vbox.SetVAlign(gtk.AlignStart) // Align vertically to the top (start)
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
