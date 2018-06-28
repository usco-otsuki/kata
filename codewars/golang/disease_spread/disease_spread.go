package disease_spread

func Epidemic(tm, n, s0, i0 int, b, a float64) int {
	s := make([]float64, 2)
	i := make([]float64, 2)
	r := make([]float64, 2)
	dt := float64(tm) / float64(n)
	max := 1.0
	s[0] = float64(s0)
	i[0] = float64(i0)
	for k := 0; k < n; k++ {
		p := (k + 1) % 2
		q := k % 2
		s[p] = s[q] - dt*b*s[q]*i[q]
		i[p] = i[q] + dt*(b*s[q]*i[q]-a*i[q])
		r[p] = r[q] + dt*i[q]*a
		if max < i[p] {
			max = i[p]
		}
	}

	return int(max)
}
