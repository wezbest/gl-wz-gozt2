/*
Refers to the other packages
*/

package main

import (
	"pkg/display"
	"pkg/msg"
)

func main() {
	display.BanShow()
	msg.Hi()
	display.Display("fuck you")
	msg.Exciting("SmellHerAss")
}
