package josephus_permutation

func Josephus(items []interface{}, k int) []interface{} {
	res := make([]interface{}, 0, len(items))

	curLen := len(items)

	for i := k - 1; curLen > 0; {
		for ; i >= curLen; {
			i -= curLen
		}

		res = append(res, items[i])
		items = append(items[:i], items[i+1:]...)
		curLen--
		i = i - 1 + k
	}

	return res
}