package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) getPriceText() (open, current, change *canvas.Text) {
	var g Gold

	gold, err := g.GetPrices()
	if err != nil {
		app.Errorlog.Println(err)

		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		open = canvas.NewText("open: Unreachable", grey)
		current = canvas.NewText("current: Unreachable", grey)
		change = canvas.NewText("change: Unreachable", grey)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
		}

		openTxt := fmt.Sprintf("open: %.3f %s", gold.PreviousClose, currency)
		currentTxt := fmt.Sprintf("current: %.3f %s", gold.Price, currency)
		changeTxt := fmt.Sprintf("change: %.3f %s", gold.Change, currency)

		open = canvas.NewText(openTxt, displayColor)
		current = canvas.NewText(currentTxt, displayColor)
		change = canvas.NewText(changeTxt, displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return
}
