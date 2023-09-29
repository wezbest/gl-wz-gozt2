/*
Arrays
*/

package lib

import (
	f "fmt"
)

// structure of room
type Room struct {
	name string
	cln  bool
}

/*
In the function below there is an array which is being created inside the function
- note that it is an array of structs - [4]Room , which means that the array will have 4 elements
of type Room struct - So each element will have elemeent- name of rooma and cln ( whether it is clean
or not , an din this case it is a bool - True or False) )
*/
func chkCln(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cln {
			f.Println(room.name, " is clean")
		} else {
			f.Println(room.name, " is not clean")
		}
	}
}

// Make array of room
func M1f() {
	rooms := [...]Room{
		{name: "Off"}, // default for strucelement cln will be false in all cases
		{name: "War"},
		{name: "Rec"},
		{name: "Ops"},
	}

	/* This function defined above will just spit about the room names basically and the default
	cln will be false. Then in the next section we will defined the cln variable and run a check again
	*/
	chkCln(rooms)

	Texc("Runnning cleaning function here")
	f.Println("Performing Cleaning...")
	rooms[2].cln = true
	rooms[3].cln = true

	chkCln(rooms)

}
