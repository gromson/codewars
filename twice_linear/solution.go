package twice_linear

import "sort"

// [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
// 0. 3
// 1. 7
// 2. 15
// 3. 31
// 4. 63

type Queue []int

func (q *Queue) push(i ...int) {
	*q = append(*q, i...)
	sort.Ints(*q)
}

func (q *Queue) pop() int {
	r := (*q)[0]
	*q = (*q)[1:]
	return r
}

func DblLinear(n int) int {
	u := buildSequence(n)
	return u[n]
}

func buildSequence(n int) []int {
	u := []int{1}
	q := make(Queue, 0)

	y, z := getYZ(1)
	q.push(y, z)

	for len(u) <= n {
		x := q.pop()

		if u[len(u)-1] == x {
			continue
		}

		y, z := addValue(&u, x)

		//if len(u) <= n {
			q.push(y, z)
		//}
	}

	//steps := int(math.Ceil(math.Log2(float64(n)))) + 1
	//
	//for s := 0; s <= steps; s++ {
	//	y1, z1, y2, z2 := addValues(&u, y, z, n)
	//	addValues(&u, y1, y2, n)
	//	y, z = z1, z2
	//}

	return u
}

func getYZ(x int) (y, z int) {
	y = x*2 + 1
	z = x*3 + 1
	return
}

func addValues(u *[]int, x1, x2 int) (y1, z1, y2, z2 int) {
	//if len(*u) <= n {
	*u = append(*u, x1, x2)
	y1, z1 = getYZ(x1)
	y2, z2 = getYZ(x2)
	//}

	return
}

func addValue(u *[]int, x int) (y, z int) {
	*u = append(*u, x)
	y, z = getYZ(x)
	return
}
