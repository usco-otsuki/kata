package easy_diagonal

func Diagonal(n, p int) int {
	dp := [][]int{}
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, p+1))
		dp[i][0] = i + 1
	}
	for j := 1; j <= p; j++ {
		for i := j; i <= n; i++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp[n][p]
}
