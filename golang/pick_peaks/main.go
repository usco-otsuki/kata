package pick_peaks

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	if len(array) == 0 {
		return PosPeaks{[]int{}, []int{}}
	}
	pos := []int{}
	peaks := []int{}
	prev := array[0]
	riseIndex := -1
	for i := 1; i < len(array); i++ {
		if prev < array[i] {
			riseIndex = i
		} else if prev > array[i] {
			if riseIndex >= 0 {
				pos = append(pos, riseIndex)
				peaks = append(peaks, array[riseIndex])
				riseIndex = -1
			}
		}
		prev = array[i]
	}
	return PosPeaks{pos, peaks}
}
