package kata

func SplitStrings(str string) []string {
	array := []string{}
	tmp := ""
	for i := 0; i < len(str); i++ {
		if len(tmp) <= 1 {
			tmp += string(str[i])
		}
		if len(tmp) == 2 {
			array = append(array, tmp)
			tmp = ""
		}
	}
	if len(tmp) == 1 {
		array = append(array, tmp+"_")
	}

	return array
}
