package twice_linear

import "math"

func DblLinear(n int) int {
	u := buildSequence(n)
	return u[n]
}

func buildSequence(n int) []int {
	u := make([]int, 0)

	steps := int(math.Ceil(math.Log2(float64(n)))) + 1

	for s := 0; s < steps; s++ {

	}

	return u
}

func addValues(u *[]int, x, n int) {
	y, z := addValue(u, x)
	addValues(u, y, n)
	addValues(u, z, n)
}

func addValue(u *[]int, x int) (y, z int) {
	y = x*2 + 1
	z = x*3 + 1

	*u = append(*u, x, y, z)

	return
}
