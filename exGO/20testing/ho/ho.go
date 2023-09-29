/*
Writing Tests in go
FIFO Queue- First Data in and first data out struture
*/

package ho

import (
	L "20testing/lib"
	f "fmt"
)

// Defining the data structure here
type Que struct {
	items    []int
	capacity int
}

func Meho() {
	f.Println("Hello World")
	L.T1("Hello World")
}

// New function which specifies capacity and returns a Que
func New(capacity int) Que {
	return Que{make([]int, 0, capacity), capacity} // Creating a struct
}

// Create append function to add to the que - Receiver function
/*
- Making a receiver function called Append , which can be called aginast pointer Que - So
  - It can be written as q.Append(1)
  - Also checking if the que is full with the if len(q.items) == q.capacity
	- If it is full then return false
	- if it is empty then return true
*/
func (q *Que) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}
	q.items = append(q.items, item)
	return true
}

// Nxt function also a reciver = Retrieve item at the next of the Que
func (q *Que) Next() (int, bool) {
	if len(q.items) > 0 {
		next := q.items[0]
		q.items = q.items[1:]
		return next, true
	} else {
		return 0, false
	}

}
