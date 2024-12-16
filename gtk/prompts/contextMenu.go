package prompts

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	"github.com/brycensranch/rokon/core"
)

func CreateMenu(window *gtk.ApplicationWindow, app *gtk.Application) *gio.Menu {

	menu := gio.NewMenu()

	exampleMenu := gio.NewMenuItem("Example", "example")
	exampleSubMenu := gio.NewMenu()

	aboutMenuItem := gio.NewMenuItem("About This App", "app.about")
	aboutMenuItem.Connect("activate", func() {
		showAboutWindow(window, app)
	})

	updateMenuItem := gio.NewMenuItem("Check For Updates", "app.check-for-updates")

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

	checkForUpdatesAction := gio.NewSimpleAction("check-for-updates", nil)
	checkForUpdatesAction.Connect("activate", func() {
		core.CheckForUpdates()
	})

	app.AddAction(checkForUpdatesAction)
	exampleSubMenu.AppendItem(updateMenuItem)

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
