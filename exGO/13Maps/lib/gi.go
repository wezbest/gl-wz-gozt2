/*
Given solution from q.txt
*/

package lib

import f "fmt"

const (
	gOnline      = 0
	gOffline     = 1
	gMaintenance = 2
	gRetired     = 3
)

func G1m() {
	gservers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	// here a map is being generated from the above slice 
	gserverStatus := make(map[string]int)
	
	for _, server := range gservers {
		gserverStatus[server] = gOnline
	}

	// Display server info
	Texc("Display the initial server status with no changes peformed")
	gSt(gserverStatus)

	// Change DarkStart to retired 
	gserverStatus["darkstar"] = gRetired

	// Change Aiur to offline 
	gserverStatus["aiur"] = gOffline

	// Display server info
	Texc("Display the server status after changes")
	gSt(gserverStatus)

	// Change all servers to Maintenance
	for gservers := range gserverStatus {
		gserverStatus[gservers] = gMaintenance
	}

	// Display server info
	Texc("Display the server status after changes")
	gSt(gserverStatus)

}

// Function to print server status
func gSt(servers map[string]int) {
	//total number of servers
	f.Println("\n Thre are :", len(servers), "Servers") //len(servers) gets the number of servers 

	//number of servers for each status
	gstats := make(map[int]int) // This new map - Status Code, Number of servers with the status code 
	for _, status := range servers {
		switch status {
		case gOnline:
			gstats[gOnline] += 1
		case gOffline:
			gstats[gOffline] += 1
		case gMaintenance:
			gstats[gMaintenance] += 1
		case gRetired:
			gstats[gRetired] += 1
		default:
			panic("Fucker Not good ")
		}
	}

	Texc("Printing the server status")
	// Printing information of the servers
	f.Println(gstats[gOnline], "Servers are online")
	f.Println(gstats[gOffline], "Servers are Offline")
	f.Println(gstats[gMaintenance], "Servers are Maintenance")
	f.Println(gstats[gRetired], "Servers are Retired")
}
