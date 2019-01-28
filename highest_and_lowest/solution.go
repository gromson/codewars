package highest_and_lowest

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow1(in string) string {
	md := make([]int, 0)
	sl := strings.Split(in, " ")
	for _, s := range sl {
		n, _ := strconv.Atoi(s)
		md = append(md, n)
	}
	sort.Ints(md)
	return fmt.Sprintf("%d %d", md[len(sl)-1], md[0])
}

func HighAndLow(in string) string {
	min, max := 1<<31-1, -1<<31
	sl := strings.Split(in, " ")

	for _, s := range sl {
		n, _ := strconv.Atoi(s)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
