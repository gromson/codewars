package highest_and_lowest

import "testing"

var testCases = [][]string{
	{"3 9 5 6 -5 7 34 -10", "34 -10"},
	{"23 3 5 7 8 -3 4 -6", "23 -6"},
}

func TestHighAndLow(t *testing.T) {
	for _, tc := range testCases {
		res := HighAndLow(tc[0])
		if res != tc[1] {
			t.Errorf("Result is \"%s\", expected \"%s\"", res, tc[1])
		}
	}
}

func BenchmarkHighAndLow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HighAndLow(testCases[0][0])
	}
}

func BenchmarkHighAndLow1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HighAndLow1(testCases[0][0])
	}
}
