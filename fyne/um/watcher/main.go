package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"watcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	Errorlog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HTTPClient          *http.Client
}

func main() {
	var myApp Config

	fyneApp := app.NewWithID("jp.mydns.kokoichi.watcher.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	myApp.InfoLog = log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
	myApp.Errorlog = log.New(os.Stdout, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := myApp.connectSQL()
	if err != nil {
		log.Fatal(err)
	}

	myApp.setupDB(db)

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

func (a *Config) connectSQL() (*sql.DB, error) {
	path := a.App.Storage().RootURI().Path() + "/sql.db"
	a.InfoLog.Println("path: ", path)

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (a *Config) setupDB(db *sql.DB) {
	a.DB = repository.NewSQLiteRepository(db)

	if err := a.DB.Migrate(); err != nil {
		a.Errorlog.Println(err)
		log.Fatal(err)
	}
}
