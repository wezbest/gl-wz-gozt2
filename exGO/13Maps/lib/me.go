/*
My attempt at exercise
Server Status


*/

package lib

import f "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func Mem() { // Mem = Main function
	mserverPrint()
	mserStatus()
	mserMap()
	msStatUp()
}

// - Number of servers
func mserverPrint() {
	Texc("Names of Servers")
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	f.Println("Servers : ", servers)
}

// - Number of servers for each status (Online, Offline, Maintenance, Retired)
func mserStatus() {
	Texc("Server Status")
	f.Println("Online:", Online, "\nOffline", Offline, "\nMaintenance", Maintenance, "\nRetired", Retired)
}

// * Store the existing slice of servers in a map
func mserMap() {

	// Make a map for the servers
	serverMap := make(map[string]int)
	serverMap = map[string]int{
		"darkstar": 0,
		"aiur":     0,
		"omicron":  0,
		"w359":     0,
		"baseline": 0,
	}
	//  - display server info
	Texc("Printing the server map which has been created: ")
	f.Println(serverMap)
}

func msStatUp() {
	serverMap := map[string]int{
		"darkstar": 0,
		"aiur":     0,
		"omicron":  0,
		"w359":     0,
		"baseline": 0,
	}
	Texc("Changing the status of the servers")
	//  - change `darkstar` to `Retired`
	Para("Changing darkstart to retired")
	serverMap["darkstar"] = 3
	//  - change `aiur` to `Offline`
	Para("Changing aiur to offline")
	serverMap["aiur"] = 1
	//  - display server info
	Para("Displaying server info")
	f.Println(serverMap)
	//  - change all servers to `Maintenance`
	Para("Changing all servers to maintenance")
	for k := range serverMap {
		serverMap[k] = 3
	}
	//  - display server info
	f.Println(serverMap)
	
}
