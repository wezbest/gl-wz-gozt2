/*
This go file is just for demonstrating the colors usage, which you can import as and when you need it 
*/

package lib

import (
	f "fmt"
	C "github.com/fatih/color" 
)

// The following function shouldnt be used anywhere ,its just for illustration of using colors
func FaihColorsShow() {

	// Defining all the fatih colors 
	cr := C.New(C.FgRed).SprintFunc()
	chr := C.New(C.FgHiRed).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
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

	f.Printf(cr("This Red 			:%v\n"), cr("Red"))
	f.Printf(chr("This HIRed 		:%v\n"), chr("HiRed"))
	f.Printf(cg("This Green  		:%v\n"), cg("Green"))
	f.Printf(chg("This HiGreen  	:%v\n"), chg("HiGreen"))
	f.Printf(cy("This Yellow 		:%v\n"), cy("Yellow"))
	f.Printf(chy("This HiYellow 	:%v\n"), chy("HiYellow"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	f.Printf(cr("This Red Color %v\n"), cr("Red"))
	
}