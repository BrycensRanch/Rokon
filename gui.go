package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/go-resty/resty/v2"
	"github.com/koron/go-ssdp"
	"github.com/spf13/viper"
)

func indicateScanning(window *gtk.ApplicationWindow) {
	window.SetTitle("Rokon: Control your Roku from your desktop")
	window.SetChild(&gtk.NewLabel("Searching for Rokus on your network...").Widget)
}

func fetchRokuInfo(rokuLocation string) Root {
	var root Root
	client := resty.New()
	resp, err := client.R().
		SetResult(&root).
		Get(rokuLocation)

	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Body:", resp)
	}
	return root
}

func displayDiscoveredRokus(app *gtk.Application, window *gtk.ApplicationWindow, discoveredRokus []ssdp.Service) {
	glib.IdleAdd(func() {
		if discoveredRokus == nil {
			displayRokuIPPrompt(window)
			return
		}
		log.Println("Discovered Rokus:", discoveredRokus)
		root := fetchRokuInfo(discoveredRokus[0].Location + "/")
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
	logoPath, err := findLogoFilePath()
	if err != nil {
		panic("ABORT ABORT")
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

func initializeApplicationWindow(app *gtk.Application) *gtk.ApplicationWindow {
	window := gtk.NewApplicationWindow(app)
	// Create the main menu
	menu := createMenu(window, app)
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
	logNetworkInterfaces()
	window := initializeApplicationWindow(app)
	monitorForUserActivity(window)

	// window.Maximize()
	scanOnStartup := viper.GetBool("scanOnStartup")
	// Start searching for Rokus when the app is activated
	var rokuChan chan []ssdp.Service
	if scanOnStartup {
		rokuChan = searchForRokus()
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
			return
		}
		var rokuList []string
		const (
			MaxValue = int(2)
		)
		for i, roku := range discoveredRokus {
			deviceIP, _ := getHostFromLocation(roku.Location)
			if i > 0 {
				rokuList = append(rokuList, fmt.Sprintf("Device %d: %v", i+1, deviceIP))
			} else {
				rokuList = append(rokuList, fmt.Sprintf("Device at %v", deviceIP))
			}
		}
		if len(discoveredRokus) > MaxValue {
			rokuList = append(rokuList, fmt.Sprintf("...and %d more devices", len(discoveredRokus)-MaxValue))
		}
		rokuListString := strings.Join(rokuList, "\n")

		url := discoveredRokus[0].Location + "/device-image.png"
		deviceImagePath, err := fetchImageAndDownload(url)
		if err != nil {
			log.Println("Error getting image from URL:", err)
			return
		}
		sendGTKNotification("Roku Discovered", rokuListString, deviceImagePath, app)
		displayDiscoveredRokus(app, window, discoveredRokus)
	}()
}
