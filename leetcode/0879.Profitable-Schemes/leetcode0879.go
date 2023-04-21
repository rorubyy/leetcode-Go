package leetcode

func maxVal(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	C, ans := len(group), 0
	var MOD int = 1e9 + 7
	// make a 3D array with variable
	dp := make([][][]int, C+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	// dp
	dp[0][0][0] = 1
	for i := 1; i <= C; i++ {
		g := group[i-1]
		p := profit[i-1]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j >= g {
					dp[i][j][k] = (dp[i-1][j][k] + dp[i-1][j-g][maxVal(0, k-p)]) % MOD
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}
	for i := 0; i <= n; i++ {
		ans = (ans + dp[C][i][minProfit]) % MOD
	}
	return ans
}
