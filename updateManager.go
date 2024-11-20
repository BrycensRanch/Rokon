package main

import (
	"log"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"golang.org/x/mod/semver"
)

func checkForUpdates(app *gtk.Application) {
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
}
