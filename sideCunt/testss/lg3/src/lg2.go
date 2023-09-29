// Lip GLoss script in here

package src

import (
	f "fmt"

	"github.com/charmbracelet/lipgloss"
)

// Constants being define for width and height

//Variable Definitions

var (

	s = "fofofofo"
	// Dialog Box Style
	dialogBoxStyle2 = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Background(lipgloss.Color("#874BFD")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Width(100).
		Align(lipgloss.Center).
		SetString(s)
)

// Functions will go in here

func Dia2() {

	f.Println(dialogBoxStyle2.Render())
}

// Test function to stop from showing erros all the time
func TestPrint2() {

	f.Println("Hello, World!")
}
