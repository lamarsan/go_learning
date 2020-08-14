package first_month

/**
 * description 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 *
 * @author lamar
 * @date 2020/7/30 9:41 上午
 */
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			temp := max2(j*dp[i-j], j*(i-j))
			if temp > max {
				max = temp
			}
		}
		dp[i] = max
	}
	return dp[n]
}

func max2(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
