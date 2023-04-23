package main

import (
	"fmt"
	"strconv"
	"watcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (a *Config) holdingsTab() *fyne.Container {
	a.HoldingsTable = a.getHoldingsTable()

	holdingsContainer := container.NewVBox(a.HoldingsTable)

	return holdingsContainer
}

func (a *Config) getHoldingsTable() *widget.Table {
	data := a.getHoldingsSlice()
	a.Holdings = data

	t := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(id widget.TableCellID, c fyne.CanvasObject) {
			// c.(*widget.Label).SetText(data[id.Row][id.Col].(string))
			if id.Col == (len(data[0])-1) && id.Row != 0 {
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete", "Are you sure you want to delete this holding?", func(deleted bool) {
						id, _ := strconv.Atoi(data[id.Row][0].(string))
						if err := a.DB.DeleteHolding(int64(id)); err != nil {
							a.Errorlog.Println(err)
						}

						a.refreshHoldingsTable()
					}, a.MainWindow)
				})

				w.Importance = widget.HighImportance

				c.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				c.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[id.Row][id.Col].(string)),
				}
			}
		},
	)

	colWidths := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func (a *Config) getHoldingsSlice() [][]interface{} {
	var slice [][]interface{}

	hs, err := a.currentHoldings()
	if err != nil {
		a.Errorlog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Delete?"})

	for _, h := range hs {
		var currentRow []interface{}
		currentRow = append(currentRow, strconv.FormatInt(h.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", h.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%2f", float32(h.PurchasePrice/100)))
		currentRow = append(currentRow, h.Purchased.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (a *Config) currentHoldings() ([]repository.Holdings, error) {
	h, err := a.DB.AllHoldings()
	if err != nil {
		a.Errorlog.Println(err)
		return nil, fmt.Errorf("failed to get holdings: %w", err)
	}

	return h, nil
}
