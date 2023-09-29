// Lip GLoss script in here

package src

import (
	f "fmt"

	"github.com/charmbracelet/lipgloss"
)

//Variable Definitions

var (

	// Dialog Box Style
	dialogBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		SetString("Fucker").
		Width(100).
		Align(lipgloss.Center)
)

// Functions will go in here

func Dia1() {
	f.Println(dialogBoxStyle.Render())
}

// Test function to stop from showing erros all the time
func TestPrint() {
	f.Println("Hello, World!")
}
