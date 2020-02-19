package main

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if c_lib <= c_road {
		return int64(c_lib * n)
	}

	alist := make(map[int32][]int32)
	for i := int32(1); i <= n; i++ {
		alist[i] = make([]int32, 0)
	}

	for _, road := range cities {
		alist[road[0]] = append(alist[road[0]], road[1])
		alist[road[1]] = append(alist[road[1]], road[0])
	}

	cost := int64(0)
	for i := int32(1); i <= n; i++ {
		subgraph := make([]int32, 0, 2)
		subgraph = fillSubgraph(subgraph, alist, i)
		if len(subgraph) > 0 {
			cost += int64(len(subgraph)-1)*int64(c_road) + int64(c_lib)
		}
	}

	return cost
}

func fillSubgraph(subgraph []int32, alist map[int32][]int32, city int32) []int32 {
	if _, ok := alist[city]; !ok {
		return subgraph
	}
	subgraph = append(subgraph, city)
	clist := alist[city]
	delete(alist, city)
	for _, c := range clist {
		subgraph = fillSubgraph(subgraph, alist, c)
	}
	return subgraph
}

func main() {
	res := roadsAndLibraries(5, 6, 1, [][]int32{{1, 2}, {1, 3}, {1, 4}})
	println(res)
}
