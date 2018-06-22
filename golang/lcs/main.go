package lcs

func max(strings ...string) string {
	maxString := ""
	for _, str := range strings {
		if len(str) > len(maxString) {
			maxString = str
		}
	}
	return maxString
}

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	dp := make([][]string, 2)
	dp[0] = make([]string, len(s2)+1)
	dp[1] = make([]string, len(s2)+1)

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
			if s1[i-1] == s2[j-1] {
				dp[i%2][j] = max(dp[i%2][j], dp[(i-1)%2][j-1]+string(s1[i-1]))
			}
		}
	}
	return dp[len(s1)%2][len(s2)]
}
