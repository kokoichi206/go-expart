package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// at least 1 window
	w := a.NewWindow("Hello")

	w.SetContent(widget.NewLabel("Hello Fyne!"))
	// start event loop
	w.ShowAndRun()
	// w.Show()
	// a.Run()

	something()
}

func something() {
	fmt.Println("do something")
}
