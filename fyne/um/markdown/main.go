package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// create a app
	a := app.New()

	// window
	win := a.NewWindow("Markdown")

	// get the user interface
	edit, preview := cfg.makeUI()

	// content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show and run
	win.Resize(fyne.Size{Width: 800, Height: 600})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (cfg *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	// なんじゃこれ、えぐい
	// it's not perfect
	preview := widget.NewRichTextFromMarkdown("")

	cfg.EditWidget = edit
	cfg.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}
