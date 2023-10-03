/*
This go file is just for demonstrating the colors usage, which you can import as and when you need it
*/

package lib

import (
	f "fmt"

	C "github.com/fatih/color"
)

// The following function shouldnt be used anywhere ,its just for illustration of using colors
func FatihColorsShow() {

	// Defining all the fatih colors
	cr := C.New(C.FgRed).SprintFunc()
	chr := C.New(C.FgHiRed).SprintFunc()
	chrb := C.New(C.FgHiRed, C.Bold).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
	chgbu := C.New(C.FgHiGreen, C.Bold, C.Underline).SprintFunc()
	cy := C.New(C.FgYellow).SprintFunc()
	chy := C.New(C.FgHiYellow).SprintFunc()
	cb := C.New(C.FgBlue).SprintFunc()
	chb := C.New(C.FgHiBlue).SprintFunc()
	cm := C.New(C.FgMagenta).SprintFunc()
	chm := C.New(C.FgHiMagenta).SprintFunc()
	ccy := C.New(C.FgCyan).SprintFunc()
	chcy := C.New(C.FgHiCyan).SprintFunc()
	cw := C.New(C.FgWhite).SprintFunc()
	chw := C.New(C.FgHiWhite).SprintFunc()

	// Just print using all to remove errors shit
	// There is also bold and italic and you cna declare background also

	f.Printf(cr("This Red 			:%v\n"), cr("Red"))
	f.Printf(chr("This HIRed 		:%v\n"), chr("HiRed"))
	f.Printf(chrb("This HIRed 		:%v\n"), chrb("HiRedBold"))
	f.Printf(cg("This Green  		:%v\n"), cg("Green"))
	f.Printf(chg("This HiGreen  	:%v\n"), chg("HiGreen"))
	f.Printf(chgbu("This HiGreen Bold Underline  	:%v\n"), chgbu("HiGreen"))
	f.Printf(cy("This Yellow 		:%v\n"), cy("Yellow"))
	f.Printf(chy("This HiYellow 	:%v\n"), chy("HiYellow"))
	f.Printf(cb("This Blue 			:%v\n"), cb("Blue"))
	f.Printf(chb("This HiBlue 		:%v\n"), chb("HiBlue"))
	f.Printf(cm("This Magenta 		:%v\n"), cm("Magenta"))
	f.Printf(chm("This HiMagenta  	:%v\n"), chm("HiMagenta"))
	f.Printf(cy("This Yellow  		:%v\n"), cy("Yellow"))
	f.Printf(ccy("This Cyan 		:%v\n"), ccy("Cyan"))
	f.Printf(chcy("This HiCyan  	:%v\n"), chcy("HiCyan"))
	f.Printf(cw("This White  		:%v\n"), cw("White"))
	f.Printf(chw("This HiWhite  	:%v\n"), chw("HiWhite"))

}
