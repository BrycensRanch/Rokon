package main

import (
	"fmt"
	"log"
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
	theWindow.Present()
	dialog.Present()
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
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	return dir
}

func expandUserDirectory(path string) (string, error) {
	// Replace ~ with the user's home directory
	if path[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, path[2:]), nil
	}
	return path, nil
}

func findLogoFilePath() (string, error) {
	// List of potential paths to check based on the OS
	var potentialPaths []string

	// Start with relative path (current directory or resources directory)
	potentialPaths = append(potentialPaths, "./io.github.brycensranch.Rokon.svg")
	potentialPaths = append(potentialPaths, "share/icons/hicolor/scalable/apps/io.github.brycensranch.Rokon.svg")
	potentialPaths = append(potentialPaths, "usr/share/icons/hicolor/scalable/apps/io.github.brycensranch.Rokon.svg")
	potentialPaths = append(potentialPaths, "share/Rokon/io.github.brycensranch.Rokon.svg")
	potentialPaths = append(potentialPaths, "share/icons/io.github.brycensranch.Rokon.svg")

	// Expand user directories like ~
	for _, path := range potentialPaths {
		expandedPath, err := expandUserDirectory(path)
		if err != nil {
			return "", err
		}

		// Get the absolute path based on the current working directory (if relative)
		if !filepath.IsAbs(expandedPath) {
			expandedPath = filepath.Join(getCurrentWorkingDir(), expandedPath)
		}
		log.Printf("Checking path %s for Rokon SVG Logo", expandedPath)
		// Check if the file exists at this location
		if _, err := os.Stat(expandedPath); err == nil {
			log.Printf("HIT! found at %s", expandedPath)
			return expandedPath, nil
		}
	}

	return "", fmt.Errorf("SVG logo not found in any known location")
}
