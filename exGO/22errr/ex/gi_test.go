/*
This is for running tests on he gi.go file , which is parsing he time
*/

package ex

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		// Custom test cases using the test table above - These are being shoved into the TestParseTime Function
		{"19:00:21", true},
		{"1:3:44", true},
		{"bad", true}, // Thi is a deliberate false don  by you and it will fail
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:12", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v:%v, error should be nil", data.time, err)
		}
	}
}
