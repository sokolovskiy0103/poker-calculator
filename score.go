package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Score() *container.TabItem {
	table := widget.NewTable(func() (int, int) {
		return len(g.Rounds) + 1, len(g.Players)
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("Таблиця очок")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				o.(*widget.Label).SetText(g.Players[i.Col].Name)
				o.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
			} else if i.Row-1 < len(g.Players[i.Col].Score) {
				o.(*widget.Label).SetText(strconv.Itoa(g.Players[i.Col].Score[i.Row-1]))
			} else {
				o.(*widget.Label).SetText("-")
			}
		})

	return container.NewTabItem("Результат", table)
}
