package first_month

/**
 * description 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 * 说明：每次只能向下或者向右移动一步。
 *
 * @author lamar
 * @date 2020/7/23 11:08 上午
 */
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var dp [][]int
	dp = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		dp[i][0] = sum + grid[i][0]
		sum = dp[i][0]
	}
	sum = 0
	for j := 0; j < len(grid[0]); j++ {
		dp[0][j] = sum + grid[0][j]
		sum = dp[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
