package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App        fyne.App
	InfoLog    *log.Logger
	Errorlog   *log.Logger
	MainWindow fyne.Window
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("jp.mydns.kokoichi.watcher.preferences")
	myApp.App = fyneApp

	myApp.InfoLog = log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
	myApp.Errorlog = log.New(os.Stdout, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

	myApp.MainWindow = fyneApp.NewWindow("watcher")
	myApp.MainWindow.Resize(fyne.NewSize(300, 200))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	myApp.MainWindow.ShowAndRun()
}
