package main

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	cfg.createMenuItems(win)

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

func (cfg *config) createMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("open...", cfg.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("save", func() {})
	cfg.SaveMenuItem = saveMenuItem
	cfg.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("save as.", cfg.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (cfg *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)

				return
			}

			if reader == nil {
				// user cancelled
				return
			}
			defer reader.Close()

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, win)

				return
			}

			cfg.EditWidget.SetText(string(data))

			cfg.CurrentFile = reader.URI()

			win.SetTitle(win.Title() + " - " + reader.URI().Name())
			cfg.SaveMenuItem.Disabled = false
		}, win)

		openDialog.Show()
	}
}

func (cfg *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)

				return
			}

			if writer == nil {
				// user cancelled
				return
			}

			defer writer.Close()

			// save the file
			writer.Write([]byte(cfg.EditWidget.Text))
			cfg.CurrentFile = writer.URI()

			win.SetTitle(win.Title() + " - " + writer.URI().Name())
			cfg.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.Show()
	}
}
