package sum_by_refactors

import (
	"fmt"
	"sort"
)

func SumOfDivided(lst []int) string {
	res := ""
	factors, order := findFactors(lst)

	for _, f := range order {
		res += fmt.Sprintf("(%d %d)", f, sum(factors[f]))
	}

	return res
}

func findFactors(lst []int) (pFactors map[int]map[int]struct{}, order []int) {
	pFactors = make(map[int]map[int]struct{})

	for _, n := range lst {
		factors := findNumberFactors(n)
		for _, f := range factors {
			if _, ok := pFactors[f][n]; !ok {
				if _, ok := pFactors[f]; !ok {
					pFactors[f] = make(map[int]struct{})
					order = append(order, f)
				}
				pFactors[f][n] = struct{}{}
			}
		}
	}

	sort.Ints(order)

	return
}

func findNumberFactors(n int) []int {
	factors := make([]int, 0)

	for i := 2; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, findNumberFactors(i)...)
		}
	}

	if len(factors) == 0 {
		factors = append(factors, n)
	}

	return factors
}

func sum(nums map[int]struct{}) int {
	res := 0

	for n := range nums {
		res += n
	}

	return res
}