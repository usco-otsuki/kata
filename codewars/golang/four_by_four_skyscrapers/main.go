package four_by_four_skyscrapers

var direction = []struct {
	x, y, dx, dy int
}{
	{0, 0, 1, 0},
	{0, 1, 1, 0},
	{0, 2, 1, 0},
	{0, 3, 1, 0},
	{0, 3, 0, -1},
	{1, 3, 0, -1},
	{2, 3, 0, -1},
	{3, 3, 0, -1},
	{3, 3, -1, 0},
	{3, 2, -1, 0},
	{3, 1, -1, 0},
	{3, 0, -1, 0},
	{3, 0, 0, 1},
	{2, 0, 0, 1},
	{1, 0, 0, 1},
	{0, 0, 0, 1},
}

func Perm(a []int, i int) [][]int {
	p := [][]int{}
	if i > len(a) {
		tmp := make([]int, len(a))
		copy(tmp, a)
		return [][]int{tmp}
	}
	p = append(p, Perm(a, i+1)...)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		p = append(p, Perm(a, i+1)...)
		a[i], a[j] = a[j], a[i]
	}
	return p
}

func IsValid(conf [][]int, clues []int) bool {
	for i := 0; i < 4; i++ {
		m := map[int]bool{}
		for j := 0; j < 4; j++ {
			m[conf[j][i]] = true
		}
		if len(m) != 4 {
			return false
		}
	}
	for i, clue := range clues {
		if clue == 0 {
			continue
		}
		dir := direction[i]
		x := dir.x
		y := dir.y
		count := 1
		max := conf[x][y]
		x += dir.dx
		y += dir.dy
		for j := 0; j < 3; j++ {
			if conf[x][y] > max {
				count++
				max = conf[x][y]
			}
			x += dir.dx
			y += dir.dy
		}
		if count != clue {
			return false
		}
	}
	return true
}

func SolvePuzzle(clues []int) [][]int {
	p := Perm([]int{1, 2, 3, 4}, 0)
	for _, row1 := range p {
		for _, row2 := range p {
			for _, row3 := range p {
				for _, row4 := range p {
					conf := [][]int{row1, row2, row3, row4}
					if IsValid(conf, clues) {
						return conf
					}
				}
			}
		}
	}
	return [][]int{}
}
