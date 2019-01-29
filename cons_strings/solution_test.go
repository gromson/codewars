package cons_strings

import "testing"

var strSlice = []string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail", "erfs", "erfswesdfsf", "rfsdvs", "asdrgr", "sdlfkgkdfopmld"}

func BenchmarkLongestConsec(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongestConsec(strSlice, 3)
	}
}

func BenchmarkLongestConsec1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongestConsec1(strSlice, 3)
	}
}
