package number_of_integer_partitions

func Partitions(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 2; j <= i; j++ {
			dp[i][j] = dp[i-j][j] + dp[i-1][j-1]
		}
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += dp[n][i]
	}
	return sum
}
