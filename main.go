package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	b := a.NewWindow("Cost control")
	b.Resize(fyne.NewSize(300, 500))
	entry := widget.NewEntry()
	entry2 := widget.NewEntry()
	line := canvas.NewLine(color.Black)
	container1 := container.NewVBox(entry)
	container2 := container.NewVBox(entry2)
	container_all := container.NewBorder(nil, nil, container1, container2, line)

	b.SetContent(container.NewHBox(container_all))
	b.CenterOnScreen()
	b.ShowAndRun()
}
