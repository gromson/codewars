package main

import (
	"fmt"
	"math"
)

func Solution_(A string, B string) int {
	aMap := make(map[rune]int)
	for _, a := range A {
		if _, ok := aMap[a]; !ok {
			aMap[a] = 1
			continue
		}
		aMap[a]++
	}

	bMap := make(map[rune]int)
	for _, b := range B {
		if _, ok := bMap[b]; !ok {
			bMap[b] = 1
			continue
		}
		bMap[b]++
	}

	res := 0

	for a, na := range aMap {
		if nb, ok := bMap[a]; ok {
			res += int(math.Abs(float64(na - nb)))
			delete(bMap, a)
			continue
		}

		res += na
	}

	for _, nb := range bMap {
		res += nb
	}

	return res
}
func Solution(A string, B string) int {
	aMap := make(map[rune]int)
	for _, a := range A {
		if _, ok := aMap[a]; !ok {
			aMap[a] = 1
			continue
		}
		aMap[a]++
	}

	res := 0

	for _, b := range B {
		n, ok := aMap[b]

		if !ok {
			res++
			continue
		}

		if n > 1 {
			aMap[b] -= 1
			continue
		}

		delete(aMap, b)
	}

	for _, n := range aMap {
		res += n
	}

	return res
}

func main() {
	r := Solution("something", "somewhere")
	fmt.Println(r, r == 8)
}
