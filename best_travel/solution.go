package best_travel

/*
t - max distance
k - number of towns
 */
func ChooseBestSum(t, k int, ls []int) int {
	return check(t, k, ls, 0, 0)
}

func check_(t, k int, ls []int, start int) int {
	length := len(ls)
	max := 0
	sum := 0
	n := 0
	for i := start; i < length; i++ {
		if n < k && ls[i] <= t {
			sum += ls[i]
			n++
			t -= ls[i]
		}

		max = check_(t, k-n, ls, start+1)
	}

	if n < k {
		return -1
	}

	if sum < max {
		sum = max
	}

	if sum == 0 {
		return -1
	}

	return sum
}

func check(t, k int, ls []int, sum, start int) int {
	length := len(ls)
	max := 0
	for i := start; i < length; i++ {
		if k > 0 && ls[i] <= t {
			sum += ls[i]
			t -= ls[i]
			k--
		}

		if k == 0 && ls[i] <= t+ls[i-1] {
			newSum := sum - ls[i-1] + ls[i]
			if newSum > sum {
				sum = newSum
				t = t + ls[i-1] - ls[i]
			}
		}

		max = check(t, k, ls, sum, i+1)
	}

	if k > 0 {
		return -1
	}

	if max > sum {
		sum = max
	}

	return sum
}
