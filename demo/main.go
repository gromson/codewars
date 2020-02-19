package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)
	res := 1
	for _, n := range A {
		if n <= 0 {
			continue
		}
		if n == res {
			res++
		}
		if n > res {
			break
		}
	}

	return res
}

func main() {
	n := Solution([]int{1, 3, -2, 6, 4, 1, -4, 2, 8, 7})
	fmt.Println(n == 5)
}
