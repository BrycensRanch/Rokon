package main

import (
	"log"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"golang.org/x/mod/semver"
)

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
