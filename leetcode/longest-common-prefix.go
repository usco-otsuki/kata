package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	first := strs[0]
	restStrs := strs[1:]
	commonPrefix := ""
	for p := 0; p < len(first); p++ {
		for _, str := range restStrs {
			if p >= len(str) || str[p] != first[p] {
				return commonPrefix
			}
		}
		commonPrefix += first[p : p+1]
	}
	return commonPrefix
}

func main() {}
