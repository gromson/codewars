package best_travel

/*
t - max distance
k - number of towns
 */

func ChooseBestSum(t, k int, ls []int) int {
	n := len(ls)
	maxDistance := -1
	setLength := 1 << uint(n)

	for set := 0; set < setLength; set++ {
		curDistance := 0
		if bitCout(set) != k {
			continue
		}

		for i := 0; i < n; i++ {
			mask := 1 << uint(i)
			if set&mask > 0 {
				curDistance += ls[i]
			}
		}

		if curDistance <= t && maxDistance < curDistance {
			maxDistance = curDistance
		}
	}

	return maxDistance
}

func bitCout(n int) int {
	counter := 0
	for ; n > 0; counter++ {
		n &= n - 1
	}
	return counter
}