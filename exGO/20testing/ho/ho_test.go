/*
These are the tests which are being created for the function in ho.go
*/

package ho

import (
	"fmt"
	"testing"
)

// Test ability to add items to the queue
func TestAddQue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect Que Element: %v, want%v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Shouldnt be able to add to full que")
	}
}

// Testing he Next function which adds the next item
func TestNext(t *testing.T) {

	// Defining 3 elements here
	q := New(3)

	// Loop for adding 3 items in , instead of doing it manually which you would have done
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	// This for retrieving the next item , in the correct order
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			fmt.Printf("Total items in the que are %v", q)
			t.Errorf("Should be able to get items from queue")
		}
		if item != i {
			t.Errorf("Got item in wrong order: %v. wamt %v", item, i)
		}
	}

	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("Shouldnt be able to get items from queue")
	}
	if item != 0 {
		t.Errorf("Sholdnt be mor eitems in que, got: %v", item)
	}

}
