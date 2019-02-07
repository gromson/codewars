package buddy_pairs

import "math"

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {

	for n := start; n <= limit; n++ {

	}

	return []int{}
}

func findDividers(num int) []int {
	res := make([]int, 0)
	sqrt := int(math.Sqrt(float64(num)))

	for n := 1; n <= sqrt; n++ {
		if num%n == 0 {
			res = append(res, n)
		}
	}

	for n := sqrt; n >= 1; n-- {
		if num%n == 0 && num/n != n {
			res = append(res, num/n)
		}
	}

	return res
}
