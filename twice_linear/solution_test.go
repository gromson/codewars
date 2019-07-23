package twice_linear

import "testing"

var testCases = [][2]int{{0, 1}, {2, 4}, {5, 10}, {50, 175}, {100, 447}, {500, 3355}, {1000, 8488}}

func TestDblLinear(t *testing.T) {
	for _, test := range testCases {
		res := DblLinear(test[0])
		if res != test[1] {
			t.Errorf("key %d is expected to be %d, %d given", test[0], test[1], res)
		}
	}
}

func BenchmarkDblLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DblLinear(testCases[6][0])
	}
}