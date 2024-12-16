package core

import (
	"log"

	"github.com/go-resty/resty/v2"
	"golang.org/x/mod/semver"
)

func CheckForUpdates() {
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
		log.Printf("Update Check Failed. Unable to fetch release info: %v" , resp.Status())
		return
	}

	release := resp.Result().(*GitHubRelease)
	latestReleaseVersion := release.TagName
	if semver.Compare(latestReleaseVersion, Version) > 0 {
		log.Printf("A new version is available: %v", latestReleaseVersion)
	} else if semver.Compare(latestReleaseVersion, Version) == 0 {
		log.Println("Rokon is up to date!")
	} else {
		log.Println("You are using a newer version than the available release.")
	}
}
