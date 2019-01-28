package tortoise_racing

import (
	"testing"
)

type Case struct {
	v1  int
	v2  int
	g   int
	res [3]int
}

var testCases = []Case{
	{720, 850, 70, [3]int{0, 32, 18}},
	{820, 81, 550, [3]int{-1, -1, -1}},
	{80, 91, 37, [3]int{3, 21, 49}},
}

func TestRace(t *testing.T) {
	for _, c := range testCases {
		res := Race(c.v1, c.v2, c.g)
		if res != c.res {
			t.Errorf("Result %v, expected %v", res, c.res)
		}
	}
}

func TestRace2(t *testing.T) {
	for _, c := range testCases {
		res := Race2(c.v1, c.v2, c.g)
		if res != c.res {
			t.Errorf("Result %v, expected %v", res, c.res)
		}
	}
}

func BenchmarkRace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Race(testCases[0].v1, testCases[0].v2, testCases[0].g)
	}
}

func BenchmarkRace2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Race2(testCases[0].v1, testCases[0].v2, testCases[0].g)
	}
}
