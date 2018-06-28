package range_extraction

import (
	"strconv"
	"strings"
)

func flushRange(array []string, start, end int) []string {
	switch {
	case end-start >= 2:
		array = append(array, strconv.Itoa(start)+"-"+strconv.Itoa(end))
	case end-start == 1:
		array = append(array, strconv.Itoa(start)+","+strconv.Itoa(end))
	case end-start == 0:
		array = append(array, strconv.Itoa(start))
	default:
		// do nothing
	}
	return array
}

func Solution(list []int) string {
	result := []string{}
	isFirst := true
	start, end := 0, 0
	for _, num := range list {
		if isFirst {
			start, end = num, num
			isFirst = false
		} else {
			if end+1 == num {
				end++
			} else {
				result = flushRange(result, start, end)
				start, end = num, num
			}
		}
	}
	result = flushRange(result, start, end)

	return strings.Join(result, ",")
}
