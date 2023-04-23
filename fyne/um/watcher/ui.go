package main

import (
	"time"

	"fyne.io/fyne/v2"
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

	holdingsTabContent := app.holdingsTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)

	// update in the background
	go func() {
		for range time.Tick(30 * time.Second) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	openPrice, currentPrice, priceChange := app.getPriceText()

	app.PriceContainer.Objects = []fyne.CanvasObject{openPrice, currentPrice, priceChange}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}

func (a *Config) refreshHoldingsTable() {
	a.Holdings = a.getHoldingsSlice()
	a.HoldingsTable.Refresh()
}
