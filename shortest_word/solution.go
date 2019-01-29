package shortest_word

import (
	"math"
	"strings"
	"unicode/utf8"
)

func FindShort(s string) int {
	min := 1<<31-1
	sl := strings.Split(s, " ")
	for _, w := range sl {
		curlen := utf8.RuneCountInString(w)
		if curlen < min {
			min = curlen
		}
	}
	return min
}

func FindShort2(s string) int {
	// your code
	strLen, minLen, curLen := utf8.RuneCountInString(s), math.MaxFloat64, .0

	for i := 0; i < strLen; i++ {
		if s[i:i+1] == " " {
			minLen = math.Min(minLen, curLen)
			curLen = .0
			continue
		}
		curLen++
	}
	minLen = math.Min(minLen, curLen)
	return int(minLen)
}