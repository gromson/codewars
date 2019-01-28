package statistics

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type data struct {
	lowest  float64
	highest float64
	sum     float64
	med     []float64
}

func (d *data) prepare(val float64, force bool) {
	d.checkHighest(val, force)
	d.checkLowest(val, force)
	d.add(val)
	d.prepareMed(val)
}

func (d *data) checkLowest(val float64, force bool) {
	if val < d.lowest || force {
		d.lowest = val
	}
}

func (d *data) checkHighest(val float64, force bool) {
	if val > d.highest || force {
		d.highest = val
	}
}

func (d *data) add(val float64) {
	d.sum += val
}

func (d *data) prepareMed(val float64) {
	d.med = append(d.med, val)
}

func (d *data) rng() time.Duration {
	return time.Duration(d.highest-d.lowest) * time.Second
}

func (d *data) average() time.Duration {
	num := float64(len(d.med))
	return time.Duration(d.sum/num) * time.Second
}

func (d *data) median() time.Duration {
	num := len(d.med)
	index := num/2

	sort.Float64s(d.med)

	if num%2 != 0 {
		return time.Duration(d.med[index]) * time.Second
	}

	return time.Duration((d.med[index-1]+d.med[index])/2) * time.Second
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d|%02d|%02d", h, m, s)
}

func Stati(strg string) string {
	if strg == "" {
		return ""
	}

	indSl := strings.Split(strg, ", ")
	data := &data{
		med: make([]float64, 0),
	}

	for i, ind := range indSl {
		indDataStr := strings.Split(ind, "|")
		h, _ := strconv.Atoi(indDataStr[0])
		m, _ := strconv.Atoi(indDataStr[1])
		sf, _ := strconv.ParseFloat(indDataStr[2], 64)
		s := int(sf)
		val, _ := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", h, m, s))
		data.prepare(val.Seconds(), i == 0)
	}

	rng := data.rng()
	avg := data.average()
	med := data.median()

	return fmt.Sprintf(
		"Range: %s Average: %s Median: %s",
		fmtDuration(rng),
		fmtDuration(avg),
		fmtDuration(med),
	)

}
