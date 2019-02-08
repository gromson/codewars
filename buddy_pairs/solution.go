package buddy_pairs

import "math"

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {
	checked := make(map[int]int)

	for n := start; n <= limit; n++ {
		m, mBuddy := 0, 0

		if v, ok := checked[n]; ok {
			m = v
		} else {
			m = sumDividers(n) - 1
			checked[n] = m
		}

		if m <= n {
			continue
		}

		if v, ok := checked[m]; ok {
			mBuddy = v
		} else {
			mBuddy = sumDividers(m) - 1
			checked[m] = mBuddy
		}

		if mBuddy == n {
			return []int{n, m}
		}
	}

	return nil
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
		if num%n == 0 && num/n != n && num/n != num {
			res = append(res, num/n)
		}
	}

	return res
}

func sumDividers(n int) int {
	dividers := findDividers(n)
	res := 0
	for _, d := range dividers {
		res += d
	}

	return res
}

// ------------

func d(n int) int {
	r, i := 1, 2

	for i*i < n {
		if n%i == 0 {
			r += i + n/i
		}
		i++
	}

	if i*i == n {
		r += i
	}

	return r
}
func Buddy2(start, limit int) []int {
	for n := start; n <= limit; n++ {
		m := d(n) - 1
		if n == d(m)-1 && n < m {
			return []int{n, m}
		}
	}

	return nil
}
