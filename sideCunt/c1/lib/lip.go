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
	// Variable for the dialog box
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

	t1 = lp.NewStyle().
		Foreground(lp.Color("#16FF00"))
		// Style for T2()
	t2 = lp.NewStyle().
		Foreground(lp.Color("#FF008E"))
		// Style for T3()
	t3 = lp.NewStyle().
		Foreground(lp.Color("#00DFA2"))
)

func Dia1(s string) {
	f.Println(dialogBoxStyle.Render(s))
}

// These are the simple color strings
func T1(s string) {
	f.Println(t1.Render(s))
}

func T2(s string) {
	f.Println(t2.Render(s))
}

func T3(s string) {
	f.Println(t3.Render(s))
}
