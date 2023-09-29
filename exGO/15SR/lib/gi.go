/*
Given solution
//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "fmt"

func main() {

}

*/

package lib

import (
	f "fmt"
	"time"
)

// Begin by making some types

// book Titles
type Title string // book title
type Name string  // library member name

// Tracking checkout time and checkin time
type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

// Structure for members
/*
// This has the TITLE which is a type declared above and LendAudit which is a struct declared. So this one type has multiple types within it
- Note the struct type can also have map subtypes within it as shown in the struct below
*/
type Member struct {
	name  Name
	books map[Title]LendAudit // Name of the books and title is a map declared inside the member struct 
}

// Book Entries tracking of the book
/*
This is required for knowing the total number of books and total lent out. Initially you wouldn't have
thought about doing it this way
*/
type BookEntry struct {
	total  int // Total number of books which library has
	lended int // Total books which have been lended outs
}

// Library Structure
/*
In this lib struct
- Members is a map container the key value NAME (declared earlier as a string) and Value is a member ( which is a
struct declared above)
- You can see the data structures are being nested here one by one
*/
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

// End of types

// Writing the functions

/*
This is the main function for checking checkout of the books in the library
*/
func printMemberAudit(member *Member) {

	/*
		For loop is iterating through the member structb  , and checking if the audit checkin is zero
		if it is zero , it means that the member hasnt taken a book , then in the else stetement if the
		book has been taken , it will
	*/
	for title, audit := range member.books {
		var returnTime string //Message displayed for return audit
		if audit.checkIn.IsZero() {
			returnTime = "[Not Returned]"
		} else {
			returnTime = audit.checkIn.String()
		}
		f.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

// Print out all the audits for all the members
/*
Iterates through all the current members which will be defined in the member struct
*/
func printMemberAuditS(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

// Prints all the books and lended out
/*
- Iterating through the Library map and printing out the books,
- Above we are just running audits to see books checkout and total number of books
*/
func printLibraryBooks(libray *Library) {
	f.Println()
	for title, book := range libray.books {
		f.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	f.Println()
}

//Function to checkout a book
/*
1. Check if the book is found
2. Check if the book has not been already lent out
3. Update the book that been lended out
*/
func checkoutBook(library *Library, title Title, member *Member) bool {
	//check if the title is actually available
	book, found := library.books[title]
	if !found {
		f.Println("Book not found")
		return false
	}
	//check if the book is already checked out
	if book.lended == book.total {
		f.Println("Book is already checked out")
		return false
	}
	// Update structure of library of lended
	/*
		1. Increment the total number of books lended
		2. Then accessing the library structs , book map reassign that value back to book , which
		will then allow for checking if the book is found in the loop above
	*/
	book.lended += 1
	library.books[title] = book

	/*
		update member function
		- Also assign the checkout time to the book which has been checked out in the function below
	*/
	member.books[title] = LendAudit{
		checkOut: time.Now(),
	}
	return true
}

/*
Return Book function here
*/
func returnBook(library *Library, title Title, member *Member) bool {
	//check if the title is actually available
	book, found := library.books[title]
	if !found {
		f.Println("Book not found")
		return false
	}
	// Check if the person who is checking out the book actually checked it out
	audit, found := member.books[title]
	if !found {
		f.Println("Member didnt checkout this book")
		return false
	}

	// update the library
	/*
		Decrement the lended out boooks, reassign to library
	*/
	book.lended -= 1
	library.books[title] = book

	//Update the checkin information on the member
	audit.checkIn = time.Now()
	member.books[title] = audit // reassign audit to the member information
	return true
}

/*
MAIN FUNCTION
1. If you try to keep this at the top , it doesnt see the variables which has been declared
2. So you need to move it to the bottom of the file
3. This behavior seems diff in rust
*/

func G1m() {
	//Create a new library
	library := Library{
		books:   make(map[Title]BookEntry), // make empty map for books stored
		members: make(map[Name]Member),     // make empty map for the members information
	}

	// Add books
	/*
		In the new library struct we made above ,no we are adding books
		- when lended = 0  , that means book isnt lended out and is available
		- Here the total indicates the numbers of copies of the book available in the library
	*/
	library.books["Webapps in Fo"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Lets learn go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  4,
		lended: 0,
	}

	//Adding members
	library.members["Jayson"] = Member{
		name:  "Jayson",
		books: make(map[Title]LendAudit), // This map is to keep track of all the books being lended out
	}
	library.members["Billy"] = Member{
		name:  "Billy",
		books: make(map[Title]LendAudit), // This map is to keep track of all the books being lended out
	}
	library.members["Susanna"] = Member{
		name:  "Susanna",
		books: make(map[Title]LendAudit), // This map is to keep track of all the books being lended out
	}

	// Now start by printing out the status of the library
	f.Println("\n Initial:")
	printLibraryBooks(&library)
	printMemberAuditS(&library)
  
	//Checkout a book by a member
	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member) //try to check out book
	f.Println("\n Checout a book:")                               // print out information
	// Print out the status of the library
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAuditS(&library)
	}

	//checkin a book
	returned := returnBook(&library, "Go Bootcamp", &member) //try to return book
	f.Println("\n Checkin a book :")
	// Print out the status of the library
	if returned {
		printLibraryBooks(&library)
		printMemberAuditS(&library)
	}
}
