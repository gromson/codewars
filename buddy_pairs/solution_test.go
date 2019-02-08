package buddy_pairs

import "testing"

type findDividerStruct struct {
	num int
	res []int
}

type buddyStruct struct {
	start int
	limit int
	res   []int
}

var findDividersCases = []findDividerStruct{
	{6, []int{1, 2, 3}},
	{25, []int{1, 5}},
	{75, []int{1, 3, 5, 15, 25}},
	{48, []int{1, 2, 3, 4, 6, 8, 12, 16, 24}},
}

var buddyTestCases = []buddyStruct{
	{1071625, 1103735, []int{1081184, 1331967}},
	{57345, 90061, []int{62744, 75495}},
	{2693, 7098, []int{5775, 6128}},
	{6379, 8275, nil},
}

func TestBuddy(t *testing.T) {
	for _, c := range buddyTestCases {
		res := Buddy(c.start, c.limit)
		if !equal(res, c.res) {
			t.Errorf("result %v, expected %v\n", res, c.res)
		}
	}
}

func TestBuddy2(t *testing.T) {
	for _, c := range buddyTestCases {
		res := Buddy2(c.start, c.limit)
		if !equal(res, c.res) {
			t.Errorf("result %v, expected %v\n", res, c.res)
		}
	}
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

func BenchmarkBuddy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buddy(buddyTestCases[0].start, buddyTestCases[0].limit)
	}
}

func BenchmarkBuddy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buddy2(buddyTestCases[0].start, buddyTestCases[0].limit)
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
