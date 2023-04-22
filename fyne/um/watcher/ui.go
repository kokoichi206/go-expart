package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	app.PriceContainer = priceContent

	toolbar := app.getToolBar()
	app.ToolBar = toolbar

	priceTabContent := app.pricesTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("holdings", theme.InfoIcon(), canvas.NewText("holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)
}
