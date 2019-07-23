package twice_linear

import (
	"math"
	"sort"
)

// [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
// 0. 3
// 1. 7
// 2. 15
// 3. 31
// 4. 63

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
	slice  []int
}

func NewNode(value int, parent *Node) *Node {
	return &Node{value, parent, nil, nil, make([]int, 0)}
}

func (n *Node) Grow(deep int) {
	if deep == 0 {
		return
	}

	n.Left = NewNode(n.Value*2+1, n)
	n.Right = NewNode(n.Value*3+1, n)

	n.Left.Grow(deep - 1)
	n.Right.Grow(deep - 1)
}

func (n *Node) ToSlice() []int {
	n.iterate(&n.slice)
	return n.slice
}

func (n *Node) iterate(s *[]int) {
	if n == nil {
		return
	}

	*s = append(*s, n.Value)
	n.Left.iterate(s)
	n.Right.iterate(s)
}

func DblLinear(n int) int {
	steps := int(math.Ceil(math.Log2(float64(n))))

	tree := NewNode(1, nil)
	tree.Grow(steps)

	slice := tree.ToSlice()
	sort.Ints(slice)

	res := make([]int, 0, n)
	for _, r := range slice {
		if len(res) > 0 && res[len(res)-1] == r {
			continue
		}

		res = append(res, r)
	}

	return res[n]
}

type Queue []int

func (q *Queue) push(i ...int) {
	*q = append(*q, i...)
	//sort.Ints(*q)
}

func (q *Queue) pop() int {
	r := (*q)[0]
	*q = (*q)[1:]
	return r
}

func DblLinear_(n int) int {
	u := []int{1}
	q := make(Queue, 0)

	x := 1
	y, z := x*2+1, x*3+1
	q.push(y, z)

	for len(u) <= n*2 || len(q) > 0 {
		x = q.pop()

		if u[len(u)-1] == x {
			continue
		}

		u = append(u, x)
		y, z := x*2+1, x*3+1
		if len(u) <= n*2 {
			q.push(y, z)
		}
	}

	sort.Ints(u)

	res := make([]int, 0, n)
	for _, r := range u {
		if len(res) > 0 && res[len(res)-1] == r {
			continue
		}

		res = append(res, r)
	}

	return res[n]
}
