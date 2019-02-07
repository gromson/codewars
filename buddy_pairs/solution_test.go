package buddy_pairs

import "testing"

type findDividerStruct struct {
	num int
	res []int
}

var findDividersCases = []findDividerStruct{
	{6, []int{1, 2, 3, 6}},
	{25, []int{1, 5, 25}},
}

func Test_findDividers(t *testing.T) {
	for _, c := range findDividersCases {
		res := findDividers(c.num)
		if !equal(res, c.res) {
			t.Errorf("result %v, expected %v\n", res, c.res)
		}
	}
}

func Benchmark_findDividers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDividers(1000000000)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
