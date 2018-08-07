package reversed_string

func Solution(word string) string {
	result := ""
	for i := len(word) - 1; i >= 0; i-- {
		result += word[i : i+1]
	}
	return result
}
