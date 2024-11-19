package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/brycensranch/go-aptabase/pkg/aptabase/v1"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var aptabaseClient *aptabase.Client // Package-level variable


func doTelemetry(telemetryLogger *log.Logger, app *gtk.Application, execDir string) {
	aptabaseClient := aptabase.NewClient("A-US-0332858461", version, uint64(133), false, "")
	// aptabaseClient.Logger = telemetryLogger

	if packageFormat == "detect" {
		_, err := os.Stat("/.dockerenv")

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
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer aptabaseClient.Stop()
}

func isRunningWithFirejail() bool {
	appImage := os.Getenv("APPIMAGE")
	appDir := os.Getenv("APPDIR")
	return (appImage != "" && contains(appImage, "/run/firejail")) ||
		(appDir != "" && contains(appDir, "/run/firejail"))
}

func createEvent(eventName string, eventData map[string]interface{}) {
	event := aptabase.EventData{
		EventName: eventName,
		Props:     eventData,
	}
	aptabaseClient.TrackEvent(event)
}
