// Attempting to make a lipglo function

package lib

import (
	f "fmt"
	"time"

	lg "github.com/charmbracelet/lipgloss"
)

/*
The following fnction is super sexy
1. First of all you are making a simple function that has a receiver of type string
2. Then you are doing all the coloring with a lipgloss style
3. Note that since you put a receiver of type string, you are passing in a string
4. Its part of packge main , so you can now call it in any file amazingly
5. Then you just call it in any function ensuring that what is being passed in is a string
6. You can now call this function anywhere
sna
*/

// This is called a block creation
var (
	liq = lg.NewStyle().
		Foreground(lg.Color("#16FF00")).
		Background(lg.Color("#2D033B")).
		Width(100).Align(lg.Center).Bold(true).
		Blink(true).Border(lg.DoubleBorder(), true, true, true, true).
		PaddingTop(1).
		PaddingBottom(1).
		BorderForeground(lg.Color("#FC2947"))

	banner = `

	

╱╭━━━╮╱╭━━━╮╱╱╱╱╱╭━━━╮╱╭━╮╭━╮
╱┃╭━╮┃╱┃╭━╮┃╱╱╱╱╱┃╭━━╯╱╰╮╰╯╭╯
╱┃╰━━╮╱┃╰━╯┃╱╱╱╱╱┃╰━━╮╱╱╰╮╭╯╱
╱╰━━╮┃╱┃╭╮╭╯╱╱╱╱╱┃╭━━╯╱╱╭╯╰╮╱
╱┃╰━╯┃╱┃┃┃╰╮╱╱╱╱╱┃╰━━╮╱╭╯╭╮╰╮
╱╰━━━╯╱╰╯╰━╯╱╱╱╱╱╰━━━╯╱╰━╯╰━╯



`

	lqh = lg.NewStyle().
		Background(lg.Color("#0C1E7F")).
		Foreground(lg.Color("#FF008E")).
		Bold(true).
		Underline(true)

	texco = lg.NewStyle().
		Background(lg.Color("#3A1078")).
		Foreground(lg.Color("#00DFA2"))

	pa = lg.NewStyle().
		Background(lg.Color("#2E4F4F")).
		Foreground(lg.Color("#ACFADF"))

	ko = lg.NewStyle().
		Foreground(lg.Color("#16FF00")).
		Italic(true)

	t1 = lg.NewStyle().
		Foreground(lg.Color("#16FF00"))
	t2 = lg.NewStyle().
		Foreground(lg.Color("#FF008E"))
	t3 = lg.NewStyle().
		Foreground(lg.Color("#00DFA2")).
		Italic(true)
)

func ColoBanner(s string) {
	t := time.Now().UTC().Format("Mon Jan 02 15:04:05 UTC")

	f.Println(liq.Render(banner, t, "\n\n", s))

}

// Thus funtion is for the nice heading colors
func Headr(s string) {
	f.Println("\n")
	f.Println(lqh.Render(s), "\n")
}

// This is for the sub header
func Texc(s string) {
	f.Println(texco.Render(s))
}

// This for the paragraph
func Para(s string) {
	f.Println(pa.Render(s), "\n")
}

// This is for the code
func Kode(s string) {
	f.Println(ko.Render(s))
}

// These are the simple color strings
func T1(s string) {
	f.Println(t1.Render(s))
}

func T2(s, t, u string) {
	f.Println(t2.Render(s, t, u))
}

func T3(s string) {
	f.Println(t3.Render(s))
}
