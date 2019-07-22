package twice_linear

import "math"

// [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
// 0. 3
// 1. 7
// 2. 15
// 3. 31
// 4. 63

func DblLinear(n int) int {
	u := buildSequence(n)
	return u[n]
}

func buildSequence(n int) []int {
	u := []int{1}

	y, z := getYZ(1)

	steps := int(math.Ceil(math.Log2(float64(n)))) + 1

	for s := 0; s < steps; s++ {
		y1, z1, y2, z2 := addValue(&u, y, z, n)
		addValue(&u, y1, y2, n)
		y, z = z1, z2
	}

	return u
}

func getYZ(x int) (y, z int) {
	y = x*2 + 1
	z = x*3 + 1
	return
}

func addValue(u *[]int, x1, x2, n int) (y1, z1, y2, z2 int) {
	if len(*u) <= n {
		*u = append(*u, x1, x2)
		y1, z1 = getYZ(x1)
		y2, z2 = getYZ(x2)
	}

	return
}
