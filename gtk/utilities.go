package main

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

// Function to show a dialog with the specified title and message.
func showDialog(title, message string, app *gtk.Application) {
	theWindow := app.ActiveWindow()
	dialog := gtk.NewMessageDialog(
		theWindow,
		gtk.DialogDestroyWithParent,
		gtk.MessageError,
		gtk.ButtonsNone,
	)
	dialog.SetTitle(title)
	dialog.SetApplication(app)
	// dialog.SetModal(true)
	dialog.SetChild(gtk.NewLabelWithMnemonic(message))
	dialog.Present()
}
