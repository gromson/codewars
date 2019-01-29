package cons_strings

import "strings"

func LongestConsec(strarr []string, k int) string {
	arrlen := len(strarr)
	if arrlen == 0 || k > arrlen || k <= 0 {
		return ""
	}

	index, max, length := 0, 0, 0
	for i, str := range strarr {
		if i <= k-1 {
			length += len(str)
		}

		if i > k-1 {
			length += len(str) - len(strarr[i-k])
		}

		if length > max && i >= k-1 {
			max = length
			index = i - k + 1
		}
	}

	return strings.Join(strarr[index:index+k], "")
}

func LongestConsec1(strarr []string, k int) string {
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i:i+k][:], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}
