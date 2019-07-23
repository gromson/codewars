package twice_linear

import (
	"math"
	"sort"
)

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

func (n *Node) Build(deep int) {
	if deep == 0 {
		return
	}

	n.Left = NewNode(n.Value*2+1, n)
	n.Right = NewNode(n.Value*3+1, n)

	n.Left.Build(deep - 1)
	n.Right.Build(deep - 1)
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
	if n == 0 {
		return 1
	}

	steps := int(math.Ceil(math.Log2(float64(n))))+2

	tree := NewNode(1, nil)
	tree.Build(steps)

	slice := tree.ToSlice()
	sort.Ints(slice)

	u := make([]int, 0, n)
	for _, r := range slice {
		if len(u) > 0 && u[len(u)-1] == r {
			continue
		}

		u = append(u, r)
	}

	return u[n]
}

func DblLinear2(n int) int {

	// Code
	u := []int{1}
	i := 0
	j := 0

	var y int
	var z int

	for len(u) <= n {
		y = 2*u[i] + 1
		z = 3*u[j] + 1

		if y < z {
			u = append(u, y)
			i++
		} else if y > z {
			u = append(u, z)
			j++
		} else {
			u = append(u, y)
			i++
			j++
		}
	}
	return u[len(u)-1]

}