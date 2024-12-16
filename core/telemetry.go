package core

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/brycensranch/go-aptabase/pkg/aptabase/v1"
)

var aptabaseClient *aptabase.Client

func DoTelemetry(telemetryLogger log.Logger, execDir string) {
	aptabaseClient = aptabase.NewClient("A-US-0332858461", Version, uint64(133), false, "")
	aptabaseClient.Logger = &telemetryLogger

	if PackageFormat == "detect" {

		switch {
		case os.Getenv("APPIMAGE") != "":
			PackageFormat = "AppImage"
		case FileExists("/.dockerenv"):
			PackageFormat = "docker"
		case os.Getenv("CONTAINER") == "oci":
			PackageFormat = "docker"

		case FileExists(filepath.Join(execDir, "share", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "share", "packageFormat"))
			if err == nil {
				PackageFormat = string(content)
			}
		case FileExists(filepath.Join(execDir, "usr", "share", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "packageFormat"))
			if err == nil {
				PackageFormat = string(content)
			}
		case FileExists(filepath.Join(execDir, "packaging", "usr", "share", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "packageFormat"))
			if err == nil {
				PackageFormat = string(content)
			}
		case FileExists(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat"))
			if err == nil {
				PackageFormat = string(content)
			}
		case FileExists(filepath.Join(execDir, "usr", "local", "share", "rokon", "packageFormat")):
			content, err := os.ReadFile(filepath.Join(execDir, "usr", "share", "rokon", "packageFormat"))
			if err == nil {
				PackageFormat = string(content)
			}
		default:
			PackageFormat = "native"
		}

		telemetryLogger.Println("Detected package format:", PackageFormat)
	}
	switch runtime.GOOS {
	case "linux":
		release := GetOSRelease()
		arch := runtime.GOARCH
		desktop := os.Getenv("XDG_CURRENT_DESKTOP")
		sessionType := os.Getenv("XDG_SESSION_TYPE")

		kdeSessionVersion := ""
		if desktop == "KDE" {
			kdeSessionVersion = os.Getenv("KDE_SESSION_VERSION")
		}

		telemetryLogger.Printf("Running on Linux. Specifically: %s %s with %s %s %s and %s\n",
			release, arch, desktop, os.Getenv("DESKTOP_SESSION"), kdeSessionVersion, sessionType)

		SendEvent("linux_run", map[string]interface{}{
			"arch":           arch,
			"desktop":        desktop,
			"desktopVersion": kdeSessionVersion,
			"sessionType":    sessionType,
			"packageFormat":  PackageFormat,
		})

		container := os.Getenv("container")
		if container != "" && container == "flatpak" {
			telemetryLogger.Println("Running from a Flatpak")
			SendEvent("flatpak_run", map[string]interface{}{
				"flatpak":        container,
				"flatpakVersion": Version,
			})
		} else if snap := os.Getenv("SNAP"); snap != "" {
			telemetryLogger.Println("Running from a Snap")
			SendEvent("snap_run", map[string]interface{}{
				"snap":        snap,
				"snapVersion": Version,
			})
		} else if appImage := os.Getenv("APPIMAGE"); appImage != "" {
			telemetryLogger.Println("Running from an AppImage")
			firejail := IsRunningWithFirejail()

			if firejail {
				telemetryLogger.Println("Running from an AppImage with firejail")
			}

			SendEvent("appimage_run", map[string]interface{}{
				"appimage":           filepath.Base(appImage),
				"appimageVersion":    Version,
				"firejail":           firejail,
				"desktopIntegration": os.Getenv("DESKTOPINTEGRATION"),
			})
		} else if IsPackaged == "true" {
			telemetryLogger.Println("Running from a native package")
			SendEvent("native_run", map[string]interface{}{
				"nativeVersion": Version,
				"path":          path.Base(os.Args[0]),
				"packageFormat": PackageFormat,
			})
		}
	case "windows":
		release := GetOSRelease()
		arch := runtime.GOARCH
		telemetryLogger.Printf("Running on Windows. Specifically: %s %s\n",
			release, arch)

		if PackageFormat == "portable" {
			telemetryLogger.Println("Running from a portable executable")
		}

		SendEvent("windows_run", map[string]interface{}{
			"arch":          arch,
			"version":       Version,
			"packageFormat": PackageFormat,
		})
	case "darwin":
		release := GetOSRelease()
		arch := runtime.GOARCH
		telemetryLogger.Printf("Running on macOS. Specifically: %s %s\n",
			release, arch)

		SendEvent("macos_run", map[string]interface{}{
			"arch":          arch,
			"mas":           os.Getenv("MAS"),
			"version":       Version,
			"path":          path.Base(os.Args[0]),
			"packageFormat": PackageFormat,
		})
	default:
		telemetryLogger.Printf("Unsupported telemetry platform: %s %s %s. However, the application will continue.\n",
			runtime.GOOS, GetOSRelease(), runtime.GOARCH)
		SendEvent("unsupported_platform", map[string]interface{}{
			"platform":      runtime.GOOS,
			"arch":          runtime.GOARCH,
			"version":       Version,
			"path":          path.Base(os.Args[0]),
			"packageFormat": PackageFormat,
		})
	}
	// Flush buffered events before the program terminates.
	defer aptabaseClient.Stop()
}

func IsRunningWithFirejail() bool {
	appImage := os.Getenv("APPIMAGE")
	appDir := os.Getenv("APPDIR")
	return (appImage != "" && Contains(appImage, "/run/firejail")) ||
		(appDir != "" && Contains(appDir, "/run/firejail"))
}

func SendEvent(eventName string, eventData map[string]interface{}) {
	event := aptabase.EventData{
		EventName: eventName,
		Props:     eventData,
	}
	aptabaseClient.TrackEvent(event)
}
