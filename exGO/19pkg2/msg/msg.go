package msg

// This can be auto imported if you have the go extension 

/*
This import statement is starting with pkg , instead of the directory name because 
if you see go.mod , we named - module pkg , so whatever is this name is the one which 
you need to use in the path
*/
import "pkg/display"

func Hi() {
	display.Display("fuck you")
}
