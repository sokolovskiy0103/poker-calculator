package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loadRules() *container.Scroll {
	content, err := os.ReadFile("RULES.md")
	str := string(content)
	if err != nil {
		str = "Не вдалося завантажити правила."
	}
	text := widget.NewLabel(str)
	text.Wrapping = fyne.TextWrapBreak
	return container.NewScroll(text)
}
