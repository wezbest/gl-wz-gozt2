/*
Given Solution - taken  from problemt txt
*/

package lib

import f "fmt"

func Gi() {
	switch age := 2; {
	case age == 0:
		f.Println("newborn", "Age is:", age)
	case age >= 1 && age <= 3:
		f.Println("toddler", "Age is:", age)
	case age <= 13:
		f.Println("child", "Age is:", age)
	case age <= 18:
		f.Println("teenager", "Age is:", age)
	default:
		f.Println("adult", "Age is:", age)
	}

	ShowCode2()
}

func ShowCode2() {
	Kode(`
	import f "fmt"

	func Gi() {
		switch age := 2; {
		case age == 0:
			f.Println("newborn", "Age is:", age)
		case age >= 1 && age <= 3:
			f.Println("toddler", "Age is:", age)
		case age <= 13:
			f.Println("child", "Age is:", age)
		case age <= 18:
			f.Println("teenager", "Age is:", age)
		default:
			f.Println("adult", "Age is:", age)
		}

	ShowCode2()
	}

	`)
}
