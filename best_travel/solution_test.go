package best_travel

import "testing"

type caseStruct struct {
	max  int
	num  int
	list []int
	res  int
}

var tcases = []caseStruct{
	{163, 3, []int{50, 55, 56, 57, 58}, 163},
	{163, 3, []int{50}, -1},
	{230, 3, []int{91, 74, 73, 85, 73, 81, 87}, 228},
	{331, 2, []int{91, 74, 73, 85, 73, 81, 87}, 178},
	{331, 4, []int{91, 74, 73, 85, 73, 81, 87}, 331},
	{331, 5, []int{91, 74, 73, 85, 73, 81, 87}, -1},
}

var bitTestCases = [][2]int{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{37, 3},
	{170, 4},
}

func TestChooseBestSum(t *testing.T) {
	for _, c := range tcases {
		res := ChooseBestSum(c.max, c.num, c.list)
		if res != c.res {
			t.Errorf("Result: %d, expected %d", res, c.res)
		}
	}
}

func TestChooseBestSum2(t *testing.T) {
	for _, c := range tcases {
		res := ChooseBestSum2(c.max, c.num, c.list)
		if res != c.res {
			t.Errorf("Result: %d, expected %d", res, c.res)
		}
	}
}

func Test_bitCounter(t *testing.T) {
	for _, c := range bitTestCases {
		res := bitCout(c[0])
		if res != c[1] {
			t.Errorf("Result: %d, expected %d", res, c[1])
		}
	}
}

func BenchmarkChooseBestSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChooseBestSum(tcases[0].max, tcases[0].num, tcases[0].list)
	}
}

func BenchmarkChooseBestSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChooseBestSum2(tcases[0].max, tcases[0].num, tcases[0].list)
	}
}

func Benchmark_bitCunter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitCout(bitTestCases[5][0])
	}
}
