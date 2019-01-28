package statistics

import "testing"

var testCases = [][]string{
	{"01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17", "Range: 01|01|18 Average: 01|38|05 Median: 01|32|34"},
	{"02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41", "Range: 00|31|17 Average: 02|26|18 Median: 02|22|00"},
	{"",""},
	{"02|15|59","Range: 00|00|00 Average: 02|15|59 Median: 02|15|59"},
	{"02|15|59, ||","Range: 02|15|59 Average: 01|07|59 Median: 01|07|59"},
}

func TestStati(t *testing.T) {
	for _, tc := range testCases {
		res := Stati(tc[0])
		if res != tc[1] {
			t.Errorf("Result is \"%s\", expected \"%s\"", res, tc[1])
		}
	}
}

func TestStati2(t *testing.T) {
	for _, tc := range testCases {
		res := Stati2(tc[0])
		if res != tc[1] {
			t.Errorf("Result is \"%s\", expected \"%s\"", res, tc[1])
		}
	}
}

func BenchmarkStati(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Stati(testCases[0][0])
	}
}

func BenchmarkStati2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Stati2(testCases[0][0])
	}
}