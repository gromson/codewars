package sum_by_refactors

import "testing"

type testCase struct {
	input  []int
	result string
}

var testCases = []testCase{
	{
		[]int{12, 15},
		"(2 12)(3 27)(5 15)",
	},
	{
		[]int{15, 21, 24, 30, 45},
		"(2 54)(3 135)(5 90)(7 21)",
	},
	{
		[]int{2, 4, 6},
		"(2 12)(3 6)",
	},
}

func TestSumOfDivided(t *testing.T) {
	for _, c := range testCases {
		res := SumOfDivided(c.input)
		if res != c.result {
			t.Errorf("wrong result for %v, \"%s\" expected, \"%s\" given", c.input, c.result, res)
		}
	}
}

func BenchmarkSumOfDivided(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfDivided(testCases[1].input)
	}
}
