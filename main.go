package main

import (
	"poker/game"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

var application = app.New()
var g = game.Game{}
var menuTab *container.TabItem
var gameTab *container.TabItem
var scoreTab *container.TabItem
var rulesTab *container.TabItem
var tabs = container.NewAppTabs()
var window = application.NewWindow("Poker score calculator")

func main() {
	menuTab = Menu()
	rulesTab = container.NewTabItem("Правила", loadRules())
	tabs.SetTabLocation(container.TabLocationTop)
	tabs.Append(menuTab)
	tabs.Append(rulesTab)
	window.SetContent(tabs)
	window.Resize(fyne.NewSize(400, 300))
	window.ShowAndRun()
}

func NewGame() {
	gameTab = container.NewTabItem("Гра", BidForms())
	scoreTab = Score()
	tabs.SetItems([]*container.TabItem{menuTab, gameTab, scoreTab, rulesTab})
	tabs.Select(gameTab)
}

func GameOver() {
	dialog := dialog.NewInformation(
		"Кінець гри",
		"Переможець: "+g.GetWinner().Name,
		window,
	)
	dialog.Show()
	dialog.SetOnClosed(func() {
		tabs.Select(menuTab)
	})
}
