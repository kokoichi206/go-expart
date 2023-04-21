package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	finalContent := container.NewVBox(priceContent)

	app.MainWindow.SetContent(finalContent)
}
