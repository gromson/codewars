package tank_truck

import "math"

func TankVol(h, d, vt int) int {
	hf := float64(h)
	r := float64(d) / 2
	l := float64(vt) / (math.Pi * r * r)
	theta := 2 * math.Acos(1-hf/r)
	sqrt := math.Sqrt(hf * (2*r - hf))
	s := theta/2*r*r - (r-hf)*sqrt

	return int(l * s)
}
