package main

import (
	"poker/game"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BidForms() *fyne.Container {
	return createRoundForm(
		"Почати раунд",
		func() { gameTab.Content = TrickForms() },
		func() bool { return g.GetSumOfBids() != g.GetNoCards() },
		func(player *game.Player, score int) { player.Bid = score },
	)
}

func TrickForms() *fyne.Container {
	return createRoundForm(
		"Наступний раунд",
		func() {
			if g.NextRound() {
				gameTab.Content = BidForms()
			} else {
				GameOver()
			}
		},
		func() bool { return g.GetSumOfTriks() == g.GetNoCards() },
		func(player *game.Player, score int) { player.Trick = score },
	)
}

func createRoundForm(title string, nextAction func(), scoreValidator func() bool, scoreSetter func(player *game.Player, score int)) *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("Раунд: "+strconv.Itoa(g.Round)),
		widget.NewLabel("Роздає "+g.Dealer.Name+" по "+strconv.Itoa(g.GetNoCards())+" карт(и)"),
	)
	inputContainer := container.NewVBox()
	nextRoundButton := widget.NewButton(title, nextAction)

	for i := range g.Players {
		player := &g.Players[i]
		playerContainer := container.NewVBox(widget.NewLabel(player.Name))
		selectedScore := 0
		scoreButtons := generateButtons(g.GetNoCards(), &selectedScore, func(s int) {
			scoreSetter(player, s)
			validate(nextRoundButton, scoreValidator())
		})
		playerContainer.Add(scoreButtons)
		inputContainer.Add(playerContainer)
	}
	content.Add(inputContainer)
	content.Add(nextRoundButton)

	return content
}

func generateButtons(maxScore int, selectedScore *int, callback func(score int)) *fyne.Container {
	scoreButtonsContainer := container.New(layout.NewGridLayout(5))
	var scoreButtons []*widget.Button

	for score := 0; score <= maxScore; score++ {
		scoreButton := widget.NewButton(strconv.Itoa(score), func(s int) func() {
			return func() {
				*selectedScore = s
				for _, btn := range scoreButtons {
					if btn.Text == strconv.Itoa(s) {
						btn.Importance = widget.HighImportance
					} else {
						btn.Importance = widget.MediumImportance
					}
					btn.Refresh()
				}
				callback(s)
			}
		}(score))
		scoreButtons = append(scoreButtons, scoreButton)
		scoreButtonsContainer.Add(scoreButton)
	}

	return scoreButtonsContainer
}

func validate(button *widget.Button, isValid bool) {
	if isValid {
		button.Show()
	} else {
		button.Hide()
	}
}
