package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/brycensranch/go-aptabase/pkg/aptabase/v1"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var aptabaseClient *aptabase.Client

func doTelemetry(telemetryLogger log.Logger, app *gtk.Application, execDir string) {
	aptabaseClient = aptabase.NewClient("A-US-0332858461", version, uint64(133), false, "")
	aptabaseClient.Logger = &telemetryLogger

	if packageFormat == "detect" {

		switch {
		case os.Getenv("APPIMAGE") != "":
			packageFormat = "AppImage"
		case fileExists("/.dockerenv"):
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
			packageFormat = "native"
		}

		telemetryLogger.Println("Detected package format:", packageFormat)
	}
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

		sendEvent("linux_run", map[string]interface{}{
			"arch":           arch,
			"desktop":        desktop,
			"desktopVersion": kdeSessionVersion,
			"sessionType":    sessionType,
			"packageFormat":  packageFormat,
		})

		container := os.Getenv("container")
		if container != "" && container == "flatpak" {
			telemetryLogger.Println("Running from a Flatpak")
			sendEvent("flatpak_run", map[string]interface{}{
				"flatpak":        container,
				"flatpakVersion": version,
			})
		} else if snap := os.Getenv("SNAP"); snap != "" {
			telemetryLogger.Println("Running from a Snap")
			sendEvent("snap_run", map[string]interface{}{
				"snap":        snap,
				"snapVersion": version,
			})
		} else if appImage := os.Getenv("APPIMAGE"); appImage != "" {
			telemetryLogger.Println("Running from an AppImage")
			firejail := isRunningWithFirejail()

			if firejail {
				telemetryLogger.Println("Running from an AppImage with firejail")
			}

			sendEvent("appimage_run", map[string]interface{}{
				"appimage":           filepath.Base(appImage),
				"appimageVersion":    version,
				"firejail":           firejail,
				"desktopIntegration": os.Getenv("DESKTOPINTEGRATION"),
			})
		} else if isPackaged == "true" {
			telemetryLogger.Println("Running from a native package")
			sendEvent("native_run", map[string]interface{}{
				"nativeVersion": version,
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

		sendEvent("windows_run", map[string]interface{}{
			"arch":          arch,
			"version":       version,
			"packageFormat": packageFormat,
		})
	case "darwin":
		release := getOSRelease()
		arch := runtime.GOARCH
		telemetryLogger.Printf("Running on macOS. Specifically: %s %s\n",
			release, arch)

		sendEvent("macos_run", map[string]interface{}{
			"arch":          arch,
			"mas":           os.Getenv("MAS"),
			"version":       version,
			"path":          path.Base(os.Args[0]),
			"packageFormat": packageFormat,
		})
	default:
		telemetryLogger.Printf("Unsupported telemetry platform: %s %s %s. However, the application will continue.\n",
			runtime.GOOS, getOSRelease(), runtime.GOARCH)
		sendEvent("unsupported_platform", map[string]interface{}{
			"platform":      runtime.GOOS,
			"arch":          runtime.GOARCH,
			"version":       version,
			"path":          path.Base(os.Args[0]),
			"packageFormat": packageFormat,
		})
	}
	// Flush buffered events before the program terminates.
	defer aptabaseClient.Stop()
}

func monitorForUserActivity(window *gtk.ApplicationWindow) {
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
}

func isRunningWithFirejail() bool {
	appImage := os.Getenv("APPIMAGE")
	appDir := os.Getenv("APPDIR")
	return (appImage != "" && contains(appImage, "/run/firejail")) ||
		(appDir != "" && contains(appDir, "/run/firejail"))
}

func sendEvent(eventName string, eventData map[string]interface{}) {
	event := aptabase.EventData{
		EventName: eventName,
		Props:     eventData,
	}
	aptabaseClient.TrackEvent(event)
}
