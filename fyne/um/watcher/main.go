package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	Errorlog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	ToolBar        *widget.Toolbar
	HTTPClient     *http.Client
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("jp.mydns.kokoichi.watcher.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	myApp.InfoLog = log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
	myApp.Errorlog = log.New(os.Stdout, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

	myApp.MainWindow = fyneApp.NewWindow("watcher")
	myApp.MainWindow.Resize(fyne.NewSize(779, 420))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	ctrlW := &desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierControl}
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierSuper}
	myApp.MainWindow.Canvas().AddShortcut(ctrlW, func(shortcut fyne.Shortcut) {
		log.Println("We tapped ctrl+w")
		myApp.MainWindow.Close()
	})
	myApp.MainWindow.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Cmd+w")
		myApp.MainWindow.Close()
	})

	myApp.makeUI()

	myApp.MainWindow.ShowAndRun()
}
