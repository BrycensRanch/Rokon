package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/brycensranch/go-aptabase/pkg/osinfo/v1"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/rand"
)

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

// Helper function to check if a string contains a substring.
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
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

func fetchImageAndDownload(url string) (string, error) {
	client := resty.New()
	resp, err := client.SetOutputDirectory(tempDirPath).EnableTrace().R().
		// SetDebug(true).
		EnableTrace().
		SetOutput(filepath.Join(tempDirPath, filepath.Base(url))).
		Get(url)
	if err != nil {
		return "", err
	}
	successfulHTTPCode := 200
	if resp.StatusCode() != successfulHTTPCode {
		return "", fmt.Errorf("failed to get image: status code %d", resp.StatusCode())
	}
	imagePath := filepath.Join(tempDirPath, filepath.Base(url))

	return imagePath, nil
}

func randomString(length int) string {
	// Define the characters to choose from
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator to get different results each time
	rand.Seed(uint64(time.Now().UnixNano())) // Cast to uint64

	// Create a string builder to efficiently build the string
	var sb strings.Builder
	for i := 0; i < length; i++ {
		// Get a random index in the charset
		randomIndex := rand.Intn(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	// Return the generated string
	return sb.String()
}

// Helper function to check if a specific flag is in the args
func containsFlag(args []string, flag string) bool {
	for _, arg := range args {
		if arg == flag {
			return true
		}
	}
	return false
}

func getOSRelease() string {
	osName, osVersion := osinfo.GetOSInfo()
	return fmt.Sprintf("%s %s", osName, osVersion)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
