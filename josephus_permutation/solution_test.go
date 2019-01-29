package josephus_permutation

import "testing"

type testData struct {
	items []interface{}
	res   []interface{}
	k     int
}

var testCases = []testData{
	{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}, 2},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, []interface{}{3, 6, 2, 7, 5, 1, 4}, 3},
	{[]interface{}{}, []interface{}{}, 3},
}

func TestJosephus(t *testing.T) {
	for _, c := range testCases {
		res := Josephus(c.items, c.k)
		if !isSlisesEq(res, c.res) {
			t.Errorf("Result %v, expected %v", res, c.res)
		}
	}
}

func BenchmarkJosephus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Josephus(testCases[2].items, testCases[2].k)
	}
}

func isSlisesEq(a, b []interface{}) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
