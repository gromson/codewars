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

func ChooseBestSum2(t, k int, ls []int) int {
	outerbest := -1
	for i, d := range ls {
		// not enough remaining values for this d to work
		if len(ls) < k {
			continue
		}
		// recursively choose best from t-d, until final level k=1
		if k > 1 {
			innerbest := ChooseBestSum2(t-d, k-1, ls[i+1:])
			// if no best available at lower level, this d cant work
			if innerbest < 0 {
				continue
			}
			d += innerbest
		}
		if d <= t && d > outerbest {
			outerbest = d
		}
	}
	return outerbest
}
