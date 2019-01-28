package tortoise_racing

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	//s = v2 * t;
	//(s-g) = v1*t;
	//v2*t-g = v1*t;
	//v2*t-v1*t = g;
	//t*(v2-v1) = g;
	t := float64(g) / float64(v2-v1) * 3600
	h := int(t) / 3600
	t -= float64(h) * 3600
	m := int(t) / 60
	t -= float64(m) * 60
	return [3]int{int(h), int(m), int(t)}
}

func Race2(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int {-1, -1, -1}
	}

	t := int(float64(g) / float64(v2 - v1) * 3600.0)

	h := t / 3600
	t = t % 3600

	m := t / 60
	s := t % 60

	return [3]int {h, m, s}
}