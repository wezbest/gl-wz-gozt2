/*
This will be settings for the lipgloss stuff  from this git
https://github.com/charmbracelet/lipgloss
*/

package lib

import (
	f "fmt"

	lp "github.com/charmbracelet/lipgloss"
)

// Setting the color for border and stuff

var (
	dialogBoxStyle = lp.NewStyle().
		Border(lp.RoundedBorder()).
		BorderForeground(lp.Color("#874BFD")).
		Foreground(lp.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Width(100).
		Align(lp.Center)
)

func Dia1() {
	f.Println(dialogBoxStyle.Render("Command 1"))
}
