package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

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
