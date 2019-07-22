package best_travel

import "fmt"

/*
t - max distance
k - number of towns
 */
func ChooseBestSum(t, k int, ls []int) int {
	total := make([][]int, 0)
	res := make([]int, 0, 3)
	check(t, k, k, ls, 0, &res, &total)
	fmt.Printf("%v\n", total)

	max := -1

	for _, set := range total {
		s := sum(set)
		if s > max && s <= t {
			max = s
		}
	}
	return max
}

func check(t, k, s int, ls []int, start int, res *[]int, total *[][]int) {
	length := len(ls)
	s--
	if s < 0 {
		return
	}

	end := length - s

	for i := start; i < end; i++ {
		*res = append(*res, ls[i])

		//if len(*res) == k {
		if s==0 {
			res := make([]int,0)
			for a:=0;a<k;a++ {
				res = append(res, ls[i-a])
				//*res = []int{}
			}
			*total = append(*total, res)
		}

		if k >= 0 {
			check(t, k, s, ls, i+1, res, total)
		}
	}
}

func sum(set []int) int {
	s := 0
	for _, n := range set {
		s += n
	}
	return s
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

func check__(t, k int, ls []int, sum, start int) int {
	length := len(ls)
	max := 0
	for i := start; i < length; i++ {
		if k > 0 && ls[i] <= t {
			sum += ls[i]
			t -= ls[i]
			k--
		}

		if k > 0 && ls[i] > t {
			continue
		}

		if k == 0 && ls[i] <= t+ls[i-1] {
			newSum := sum - ls[i-1] + ls[i]
			if newSum > sum {
				sum = newSum
				t = t + ls[i-1] - ls[i]
			}
		}

		max = check__(t, k, ls, sum, i+1)
		for j := i + 1; j < length; j++ {
			res := check__(t, k, ls, sum, j)
			if res > max {
				max = res
			}
		}
	}

	if k > 0 {
		return -1
	}

	if max < sum {
		max = sum
	}

	return max
}
