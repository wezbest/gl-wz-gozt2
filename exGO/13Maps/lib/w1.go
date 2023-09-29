/*
Maps Work on go 
*/

package lib 

import (
	f "fmt"
)

// Declaring the maps
func Wo1() {
	fetishList := make(map[string]int)
	fetishList["Pus"] = 2
	fetishList["Ass"] = 2
	fetishList["Armp"] += 1
	
	// Syntax for adding another integer (value pair)
	Texc("Adding 2 extra bbw smelly ass ,mistress")
	fetishList["Ass"] += 2
	f.Println(fetishList)

	// deleting pussy fetish
	Texc("Deleting pussy juice drinking fetish")
	delete(fetishList, "Pus") // deletes all oairs 
	f.Println("PussySniffing Fetish Deleted", fetishList)

	// Find out how much ass fetish you have 
	f.Println("need", fetishList["Ass"], "Ass")

	// Check if snot fetish found 
	snot, found := fetishList["Snot"]
	f.Println("Dirty fetish wanted ? ")
	if !found {
		f.Println("nope")
	} else {
		f.Println("ask for", snot, "fetish")
	}

	// iteraing thru the map 
	Texc("Printing out fetish list")
	totalFetishes := 0
	for _, fetish := range fetishList {
		totalFetishes += fetish
		f.Println(totalFetishes)
	}
	f.Println("Mistress gives you", totalFetishes, "fetishes")

}
