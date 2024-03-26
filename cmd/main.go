package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("go-vaders")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
