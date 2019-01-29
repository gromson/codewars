package shortest_word

import "testing"

type testStruct struct {
	str string
	min int
}

var testCases = []testStruct{
	{"bitcoin take over the world maybe who knows perhaps", 3},
	{"turns out random test cases are easier than writing out basic ones", 3},
	{"lets talk about javascript the best language", 3},
	{"i want to travel the world writing code one day", 1},
	{"Lets all go on holiday somewhere very cold", 2},
}

func TestFindShort(t *testing.T) {
	for _, c := range testCases {
		res := FindShort(c.str)
		if res != c.min {
			t.Errorf("Result: %d, expected %d", res, c.min)
		}
	}
}

func TestFindShort2(t *testing.T) {
	for _, c := range testCases {
		res := FindShort2(c.str)
		if res != c.min {
			t.Errorf("Result: %d, expected %d", res, c.min)
		}
	}
}

func BenchmarkFindShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindShort(testCases[0].str)
	}
}

func BenchmarkFindShort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindShort2(testCases[0].str)
	}
}