package drying_potatoes

func Potatoes(p0, w0, p1 int) int {
	return int(float64(100-p0) / float64(100-p1) * float64(w0))
}
