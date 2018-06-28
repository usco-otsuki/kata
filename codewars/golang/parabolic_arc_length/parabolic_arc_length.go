package parabolic_arc_length

import (
	"math"
)

func LenCurve(n int) float64 {
	len := 0.0
	width := 1.0 / float64(n)
	for i := 0; i < n; i++ {
		x1 := float64(i) * width
		x2 := float64(i+1) * width
		y1 := x1 * x1
		y2 := x2 * x2
		dx := x2 - x1
		dy := y2 - y1
		len += math.Hypot(dx, dy)
	}

	return len
}
