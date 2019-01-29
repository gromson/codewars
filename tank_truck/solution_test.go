package tank_truck

import (
	"testing"
)

var testCases = [][4]int{
	{5, 7, 3848, 2940},
	{2, 7, 3848, 907},
	{3, 6, 3500, 1750},
	{3, 6, 3501, 1750},
}

func TestTankVol(t *testing.T) {
	for _, c := range testCases {
		res := TankVol(c[0], c[1], c[2])
		if res != c[3] {
			t.Errorf("Result %d, expected %d", res, c[3])
		}
	}
}

func BenchmarkTankVol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TankVol(testCases[0][0], testCases[0][1], testCases[0][2])
	}
}
