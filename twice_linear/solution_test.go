package twice_linear

import "testing"

var testCases = [][2]int{{5, 10}}

func TestDblLinear(t *testing.T) {
	for _, test := range testCases {
		res := DblLinear(test[0])
		if res != test[1] {
			t.Errorf("key %d is expected to be %d, %d given", test[0], test[1], res)
		}
	}
}
