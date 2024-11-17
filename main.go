package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	b := a.NewWindow("Cost control")
	//entry := widget.NewEntry()
	//entry2 := widget.NewEntry()

	labelWalet := widget.NewLabel("		Wallet")
	labelCard := widget.NewLabel("     			 		Card")

	//line := canvas.NewLine(color.Black)

	InputWallet := container.NewVBox(
		widget.NewEntry(),
	)
	InputCard := container.NewVBox(
		widget.NewEntry(),
	)

	containerLabel := container.NewHBox(labelWalet, labelCard)

	container_all := container.NewGridWithColumns(2,
		container.NewVBox(InputWallet),
		container.NewVBox(InputCard),
	)

	containerScroll := container.NewScroll(container_all)
	var h float32 = 100
	containerScroll.SetMinSize(fyne.NewSize(300, h))
	//containerScroll.Resize(fyne.NewSize(400, 500))

	button := widget.NewButton("+", func() {
		InputWallet.Add(widget.NewEntry())
		InputCard.Add(widget.NewEntry())
		h = h + 100
	})

	b.SetContent(container.NewVBox(container.NewVBox(containerLabel, containerScroll, button)))
	b.Resize(fyne.NewSize(500, 600))
	b.CenterOnScreen()
	b.ShowAndRun()
}
