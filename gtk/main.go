package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/brycensranch/rokon/core"
	"github.com/brycensranch/rokon/gtk/prompts"


	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/koron/go-ssdp"
	"github.com/spf13/viper"
)

func indicateScanning(window *gtk.ApplicationWindow) {
	window.SetTitle("Rokon: Control your Roku from your desktop")
	window.SetChild(&gtk.NewLabel("Searching for Rokus on your network...").Widget)
}

func main() {
	core.Main()

	switch runtime.GOOS {
	case "windows", "darwin":
		log.Println("Running on Windows or macOS.")
		// Use GLib to set the GTK_CSD environment variable for Client-Side Decorations
		glib.Setenv("GTK_CSD", "0", true)
		os.Setenv("GTK_CSD", "0")
	default:
	}
	app := gtk.NewApplication("io.github.brycensranch.Rokon", gio.ApplicationHandlesCommandLine)
	app.SetDefault()
	if core.Version != "" {
		app.SetVersion(core.Version)
	}
	log.Printf("GTK: %d.%d.%d", int(gtk.GetMajorVersion()), int(gtk.GetMinorVersion()), int(gtk.GetMicroVersion()))
	app.ConnectActivate(func() { activate(app) })
	app.ConnectCommandLine(func(commandLine *gio.ApplicationCommandLine) int {
		return activateCommandLine(app, commandLine)
	})

	// Get the command-line arguments
	args := os.Args

	// Check for the flags that should trigger empty args
	// if containsFlag(args, "--version") || containsFlag(args, "--gpplication-service") || containsFlag(args, "--help-gapplication") {
	// 	// Create an empty argument list to pass to app.Run()
	// 	args = []string{}
	// }

	if code := app.Run(args); code > 0 {
		os.Exit(code)
	}
}

func displayDiscoveredRokus(app *gtk.Application, window *gtk.ApplicationWindow, discoveredRokus []ssdp.Service) {
	glib.IdleAdd(func() {
		if discoveredRokus == nil {
			log.Println("Didn't find any Rokus on the network. Displaying Roku IP Prompt")
			displayRokuIPPrompt(window)
			return
		}
		log.Println("Discovered Rokus:", discoveredRokus)
		root := core.FetchRokuInfo(discoveredRokus[0].Location + "/")
		labelText := fmt.Sprintf("Friendly Name: %s\nIP Address: %s",
			root.Device.FriendlyName, discoveredRokus[0].Location)

		label := gtk.NewLabel(labelText)
		window.SetChild(label)
	})
}

func displayRokuIPPrompt(window *gtk.ApplicationWindow) {
	label := gtk.NewLabel("Welcome to Rokon, to get started, enter your Roku's IP address.\nTo get it's IP address, go into Settings -> Network")
	vbox := gtk.NewBox(gtk.OrientationVertical, int(7))
	vbox.SetHAlign(gtk.AlignBaselineCenter)
	vbox.SetVAlign(gtk.AlignBaselineCenter)
	logoPath, err := core.FindLogoFilePath()
	if err != nil {
		log.Println(err)
	}

	rokonLogo := gtk.NewImageFromFile(logoPath)
	rokonLogo.SetPixelSize(int(200))
	vbox.Append(rokonLogo)
	vbox.Append(&label.Widget)
	entry := gtk.NewEntry()
	entry.SetMaxLength(40)
	entry.SetMaxWidthChars(40)
	entry.SetPlaceholderText("10.0.0.123")
	entry.SetTooltipText("YOU CAN DO IT!!!")
	vbox.Append(entry)
	submitButton := gtk.NewButtonWithLabel("Submit")
	submitButton.ConnectClicked(func() {
		log.Printf("Got text: %s", entry.Text())
		ip := entry.Text()
		if core.PingIP(ip) {
			log.Printf("IP %s is reachable.\n", ip)
			core.RokuSubmit(ip)

		} else {
			log.Println("User provided incorrect IP address %v", ip)
			vbox.Append(gtk.NewLabel(fmt.Sprintf("The IP address you provided (%v) does not ping back on this computer", ip)))
			// entry.SetText("")
			entry.ErrorBell()
		}
	})
	vbox.Append(submitButton)
	window.SetChild(vbox)
	window.Focus()
	entry.GrabFocusWithoutSelecting()
}

func sendGTKNotification(title string, body string, imagePath string, app *gtk.Application) {
	notification := gio.NewNotification(title)
	notification.SetIcon(gio.NewFileIcon(gio.NewFileForPath(imagePath)))
	notification.SetDefaultAction("app.connect-roku")
	notification.SetCategory("device")
	notification.SetBody(body)

	app.SendNotification("roku-discovered", notification)
}

func displayRokuData() {
}

func activateCommandLine(app *gtk.Application, commandLine *gio.ApplicationCommandLine) int {
	args := commandLine.Arguments()
	// cobraArgs := append([]string{rootCmd.Use}, args...)
	// rootCmd.SetArgs(cobraArgs)
	core.Start(args);
	// app.SetFlags(gio.ApplicationDefaultFlags)
	app.Activate()
	return 0
}

func initializeApplicationWindow(app *gtk.Application) *gtk.ApplicationWindow {
	window := gtk.NewApplicationWindow(app)
	// Create the main menu
	menu := prompts.CreateMenu(window, app)
	app.SetMenubar(menu)
	window.SetShowMenubar(true)
	indicateScanning(window)
	windowWidth := 800
	windowHeight := 400
	window.SetDefaultSize(windowWidth, windowHeight)
	window.SetVisible(true)
	return window
}

func activate(app *gtk.Application) {
	core.LogNetworkInterfaces()
	window := initializeApplicationWindow(app)
	monitorForUserActivity(window)

	// window.Maximize()
	scanOnStartup := viper.GetBool("scanOnStartup")
	// Start searching for Rokus when the app is activated
	var rokuChan chan []ssdp.Service
	if scanOnStartup {
		rokuChan = core.SearchForRokus()
	} else {
		// Empty!
		rokuChan = make(chan []ssdp.Service)
		log.Printf("scanOnStartup is FALSE. Not scanning for Rokus.")
		displayRokuIPPrompt(window)

	}

	// Goroutine that waits for Roku discovery to finish
	go func() {
		discoveredRokus := <-rokuChan // Receive the result from the Roku discovery
		if discoveredRokus == nil {
			log.Println("GTK UI didn't receive any discoveredRokus from rokuChan")
			glib.IdleAdd(func() {
				displayRokuIPPrompt(window)
			})
			return
		}
		var rokuList []string
		const (
			MaxValue = int(2)
		)
		for i, roku := range discoveredRokus {
			deviceIP, _ := core.GetHostFromLocation(roku.Location)
			if i > 0 {
				rokuList = append(rokuList, fmt.Sprintf("Device %d: %v", i+1, deviceIP))
				continue;
			}
			rokuList = append(rokuList, fmt.Sprintf("Device at %v", deviceIP))
		}
		if len(discoveredRokus) > MaxValue {
			rokuList = append(rokuList, fmt.Sprintf("...and %d more devices", len(discoveredRokus)-MaxValue))
		}
		rokuListString := strings.Join(rokuList, "\n")

		url := discoveredRokus[0].Location + "/device-image.png"
		deviceImagePath, err := core.FetchImageAndDownload(url)
		if err != nil {
			log.Println("Error getting image from URL:", err)
			return
		}
		sendGTKNotification("Roku Discovered", rokuListString, deviceImagePath, app)
		displayDiscoveredRokus(app, window, discoveredRokus)
	}()
}
