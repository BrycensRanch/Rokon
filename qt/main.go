package main

import (
	"fmt"
	"log"
	"os"
	"github.com/brycensranch/rokon/core"

	qt "github.com/mappu/miqt/qt6"
)

func main() {
	core.Main()
	core.Start(os.Args);
	app := qt.NewQApplication(os.Args)
	app.OnDestroyed(func() {
		log.Println("NO I DONT")
	})
	qtVersion := qt.QLibraryInfo_Version()
	log.Printf("QT: %v", qtVersion.ToString())


	btn := qt.NewQPushButton3(fmt.Sprintf("Hello from QT %s", qtVersion.ToString()))
	btn.SetFixedWidth(320)

	var counter int = 0

	btn.OnPressed(func() {
		counter++
		btn.SetText(fmt.Sprintf("You have clicked the button %d time(s)", counter))
	})

	btn.Show()

	qt.QApplication_Exec()


	fmt.Println("OK!")
}
