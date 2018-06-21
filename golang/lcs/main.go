package lcs

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	p1, p2 := 0, 0
	i, j := 0, 0
	s := ""
	for p1 < len(s1) && p2 < len(s2) {
		for i := p1; i < len(s1); i++ {
			for j := p2; j < len(s2); j++ {
				if s1[i] == s2[j] {
					s += string(s1[i])
				}
			}
		}
	}
	return ""
}
