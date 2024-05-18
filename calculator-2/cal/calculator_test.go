package cal

import (
	"testing"
)

//let us now see paramitarized tests:

var testcases = []struct {
	name string
	a int
	b int
} {
	{"first test case", 10,5},
	{"second test case", 4,2},
}

func TestDivide(t *testing.T){
	for _,i := range testcases {
		t.Run(i.name, func(t *testing.T){
			expected := i.a/i.b
			got := Divide(i.a,i.b);
			if (got != expected){
				t.Errorf("Test is failing")
			}
		})
	}
}