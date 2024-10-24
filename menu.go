package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Menu() *container.TabItem {
	inputBindings := []binding.String{}
	firstBinding := binding.NewString()
	inputBindings = append(inputBindings, firstBinding)
	entryContainer := container.NewVBox(widget.NewEntryWithData(firstBinding))
	refreshEntries := func() {
		entryContainer.RemoveAll()
		for _, bind := range inputBindings {
			entryContainer.Add(widget.NewEntryWithData(bind))
		}
	}

	buttonContainer := container.New(layout.NewGridLayout(2))

	buttonContainer.Add(widget.NewButton("+", func() {
		newBinding := binding.NewString()
		inputBindings = append(inputBindings, newBinding)
		refreshEntries()
	}))

	buttonContainer.Add(widget.NewButton("-", func() {
		if len(inputBindings) > 1 {
			inputBindings = inputBindings[:len(inputBindings)-1]
			refreshEntries()
		}
	}))

	return container.NewTabItem(
		"Меню",
		container.NewVBox(
			widget.NewLabel("Введіть імена гравців:"),
			entryContainer,
			buttonContainer,
			widget.NewButton("Нова гра", func() {
				var players []string
				for i := range inputBindings {
					player, _ := inputBindings[i].Get()
					players = append(players, player)
				}
				g.InitGame(players)
				NewGame()
			}),
		))
}
