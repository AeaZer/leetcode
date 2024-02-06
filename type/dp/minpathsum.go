package dp

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch i + j {
			case 0:
				continue
			case i:
				grid[i][j] += grid[i-1][j]
			case j:
				grid[i][j] += grid[i][j-1]
			default:
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}
