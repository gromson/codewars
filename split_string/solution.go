package split_string

import (
	"regexp"
	"unicode/utf8"
)

func Solution(str string) []string {
	res := make([]string, 0)
	strlen := utf8.RuneCountInString(str)
	if strlen%2 != 0.0 {
		str += "_"
	}

	for i := 0; i < strlen; i = i + 2 {
		res = append(res, str[i:i+2])
	}

	return res
}

func SolutionRegexp(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}

func SolutionRecursion(str string) []string {
	switch {
	case len(str) == 0:
		return []string{}
	case len(str) == 1:
		return []string{str + "_"}
	case len(str) == 2:
		return []string{str}
	default:
		return append([]string{str[:2]}, Solution(str[2:])...)
	}
}

func SolutionFast(str string) []string {
	if len(str)%2 != 0 {
		str = str + "_"
	}
	llen := len(str) / 2
	bstr := []byte(str)
	sli := make([]string, llen)
	for i := 0; i < llen; i++ {
		sli[i] = string(bstr[2*i : 2*i+2])
	}
	return sli
}
