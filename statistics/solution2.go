package statistics

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func time2snd(s string) int {
	a := strings.Split(s, "|")
	h, _ := strconv.Atoi(a[0])
	mn, _ := strconv.Atoi(a[1])
	sd, _ := strconv.Atoi(a[2])
	return 3600*h + 60*mn + sd
}
func snd2time(n int) string {
	h := n / 3600
	re := n % 3600
	mn := re / 60
	s := re % 60
	return fmt.Sprintf("%02d|%02d|%02d", h, mn, s)
}
func average(a []int) int {
	var total float64 = 0
	for _, value := range a {
		total += float64(value)
	}
	return int(total / float64(len(a)))
}
func Stati2(strg string) string {
	if strg == "" {
		return ""
	}
	a := strings.Split(strg, ", ")
	t := make([]int, 0)
	for _, tt := range a {
		t = append(t, time2snd(tt))
	}
	sort.Ints(t)
	lg := len(t)
	mean := average(t)
	md := int(float64(t[(lg-1)/2]+t[lg/2]) / 2.0)
	return fmt.Sprintf("Range: %s Average: %s Median: %s", snd2time(t[lg-1]-t[0]), snd2time(mean), snd2time(md))
}
