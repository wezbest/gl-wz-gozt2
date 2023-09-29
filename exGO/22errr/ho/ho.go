/*
Demo Excercise - Creating Errors
*/

package ho

import (
	"errors"
	"fmt"
)

func Meho() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value is", value)
	}
}

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index >= len(s.values) {
		return 0, errors.New(fmt.Sprintf(".no element at index %d", index))
	} else {
		return s.values[index], nil
	}
}
