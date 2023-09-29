/*
This is going to be generated code for testing the dialog only from
lip gloss
- Code is taken from claude
*/

package lg

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 50
	height = 10
)

var (
	subtle = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#000"}

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtle)

	buttonStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Padding(1, 2)
)

func ShowDialog(question string) {

	okBtn := buttonStyle.Render("Yes")
	cancelBtn := buttonStyle.Render("No")

	buttons := lipgloss.JoinHorizontal(okBtn, cancelBtn)

	content := lipgloss.JoinVertical(
		lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(question),
		buttons,
	)

	dialog := lipgloss.Place(width, height,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(content),
	)

	fmt.Println(dialog)

}
