/*
Third attempt trying to learn that lipgloss dialog stuff
Lipgloss Import - https://github.com/charmbracelet/lipgloss
*/

package src

import (
	f "fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Var for setting the variables
var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	dialogBoxStyle3 = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)
)

func Diad3() {

	// Setting the height and width
	doc := strings.Builder{}
	
	// Dialog builder code here
	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")
	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
	dialog := lipgloss.Place(100, 3,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle3.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog + "\n\n")
	f.Println(doc.String())
}

func Brinter() {
	f.Println("Hello, World!")
}
