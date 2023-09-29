/*
17: If..Else Excercise

//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

*/

package lib

import (
	f "fmt"
)

// Days of week

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	f.Println("Granted")
}

func accessDenied() {
	f.Println("Denied")
}

// Function to determine a day
/*
1. The following function is called weekday which accepts argument day integer
2. Here the output will be a boolean ( meaning true or false )
3. Logic is saying , if the they day is less tha number 4 as described in the constants , then it should
be a weekday , since 5 and 6 are the weekend.
4. This function will be used for determining what type of day it is
*/
func weekday(day int) bool {
	return day <= 4
}

func M1Fu() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Manager
	if role == Admin || role == Manager {
		accessGranted() //* Access at any time: Admin, Manager
	} else if role == Contractor && !weekday(today) {
		accessDenied() //* Access weekends: Contractor
	} else if role == Member && weekday(today) {
		accessGranted() //* Access weekdays: Member
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted() //* Access Mondays, Wednesdays, and Fridays: Guest
	} else {
		accessDenied()
	}

	ShowCode()
}

func ShowCode() {
	Kode(`
	func weekday(day int) bool {
		return day <= 4
	}
	
	func M1Fu() {
		// The day and role. Change these to check your work.
		today, role := Tuesday, Guest
		if role == Admin || role == Manager {
			accessGranted() //* Access at any time: Admin, Manager
		} else if role == Contractor && !weekday(today) {
			accessDenied() //* Access weekends: Contractor
		} else if role == Member && weekday(today) {
			accessGranted() //* Access weekdays: Member
		} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
			accessGranted() //* Access Mondays, Wednesdays, and Fridays: Guest
		} else {
			accessDenied()
		}
	}
	
	`)
}
