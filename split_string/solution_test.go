package split_string

import "testing"

func BenchmarkSolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution("thisisthegoodstring")
	}
}

func BenchmarkSolutionRegexp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolutionRegexp("thisisthegoodstring")
	}
}

func BenchmarkSolutionRecursion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolutionRecursion("thisisthegoodstring")
	}
}

func BenchmarkSolutionFast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolutionFast("thisisthegoodstring")
	}
}