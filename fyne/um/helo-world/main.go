package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	// at least 1 window
	w := a.NewWindow("Hello")

	output, entry, btn := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.Size{Width: 500, Height: 500})

	// w.SetContent(widget.NewLabel("Hello Fyne!"))
	// start event loop
	w.ShowAndRun()
	// w.Show()
	// a.Run()

	something()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello world!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})

	// also change color
	btn.Importance = widget.HighImportance

	app.output = output

	return output, entry, btn
}

func something() {
	fmt.Println("do something")
}
