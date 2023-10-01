package display

import (
	"fmt"

	C "github.com/fatih/color"
)

// Banner variable
var s = `

███████╗ ███╗   ███╗ ███████╗ ██╗      ██╗         
██╔════╝ ████╗ ████║ ██╔════╝ ██║      ██║         
███████╗ ██╔████╔██║ █████╗   ██║      ██║         
╚════██║ ██║╚██╔╝██║ ██╔══╝   ██║      ██║         
███████║ ██║ ╚═╝ ██║ ███████╗ ███████╗ ███████╗    
╚══════╝ ╚═╝     ╚═╝ ╚══════╝ ╚══════╝ ╚══════╝    

	██████╗   █████╗  ███╗   ██╗ ████████╗ ██╗   ██╗
	██╔══██╗ ██╔══██╗ ████╗  ██║ ╚══██╔══╝ ╚██╗ ██╔╝
	██████╔╝ ███████║ ██╔██╗ ██║    ██║     ╚████╔╝ 
	██╔═══╝  ██╔══██║ ██║╚██╗██║    ██║      ╚██╔╝  
	██║      ██║  ██║ ██║ ╚████║    ██║       ██║   
	╚═╝      ╚═╝  ╚═╝ ╚═╝  ╚═══╝    ╚═╝       ╚═╝   
`

// Add color

// Capital name Display means it can be exported
func Display(msg string) {
	fmt.Println(msg)
}

func BanShow() {
	cre := C.New(C.FgHiMagenta).SprintFunc()
	fmt.Println(cre(s))
}
