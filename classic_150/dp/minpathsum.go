package dp

// 64. 最小路径和
// https://leetcode.cn/problems/minimum-path-sum/description/?envType=study-plan-v2&envId=top-interview-150

/*给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。*/

// 动态规划转移方程：
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j] // i>0 && j>0
// dp[i][j] = dp[i][j] = dp[i-1][j] + grid[i][j] // i==0 || j==0
func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][cols-1]
}
