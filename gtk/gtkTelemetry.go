package main

import (
	"log"
	"strconv"
	"github.com/brycensranch/rokon/core"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func monitorForUserActivity(window *gtk.ApplicationWindow) {
	if (!core.Telemetry) { return; }
	// Event controller setup
	keyController := gtk.NewEventControllerKey()
	keyController.SetName("keyController")
	window.AddController(keyController)

	keyController.Connect("key-pressed", func(controller *gtk.EventControllerKey, code uint) {
		log.Println(controller.Name() + " " + strconv.FormatUint(uint64(code), 10))
		const (
			RightClickCode = uint(93) // Code representing a right-click
		)
		if code == RightClickCode {
			log.Println("right clicked")
		}
	})

	focusController := gtk.NewEventControllerFocus()
	focusController.SetName("focusController")

	focusController.Connect("enter", func() {
		log.Println("Keyboard focus entered!")
	})
	window.AddController(focusController)

	gestureClick := gtk.NewGestureClick()
	gestureClick.SetName("gestureClick")
	gestureClick.Connect("pressed", func(_, numberOfPresses uint) {
		log.Printf("Number of presses %v", numberOfPresses)
	})
	window.AddController(gestureClick)
}
