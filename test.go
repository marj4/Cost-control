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

	// Заголовки колонок
	labelWallet := widget.NewLabel("Wallet")
	labelCard := widget.NewLabel("Card")

	// Поля ввода для кошелька
	entryWallet1 := widget.NewEntry()
	entryWallet2 := widget.NewEntry()
	entryWallet3 := widget.NewEntry()

	// Поля ввода для карты
	entryCard1 := widget.NewEntry()
	entryCard2 := widget.NewEntry()
	entryCard3 := widget.NewEntry()

	// Устанавливаем минимальный размер для всех полей ввода
	setEntryMinSize := func(entries ...*widget.Entry) {
		for _, entry := range entries {
			entry.SetMinSize(fyne.NewSize(200, 40))
		}
	}
	setEntryMinSize(entryWallet1, entryWallet2, entryWallet3, entryCard1, entryCard2, entryCard3)

	// Контейнеры для каждой колонки
	InputWallet := container.NewVBox(
		entryWallet1,
		entryWallet2,
		entryWallet3,
	)

	InputCard := container.NewVBox(
		entryCard1,
		entryCard2,
		entryCard3,
	)

	// Главный контейнер для двух колонок
	containerAll := container.NewHBox(
		container.NewVBox(labelWallet, InputWallet),
		container.NewVBox(labelCard, InputCard),
	)

	// Увеличиваем пространство для всего контейнера
	b.SetContent(container.NewVBox(containerAll))
	b.Resize(fyne.NewSize(600, 400)) // Достаточно места для всех элементов
	b.CenterOnScreen()
	b.ShowAndRun()
}
