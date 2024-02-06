package dp

const (
	pointDefault = iota
	pointStone
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == pointStone {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == pointStone {
			break
		}
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == pointStone {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+j == 0 {
				obstacleGrid[0][0] = 1
				continue
			}
			if i == 0 && j > 0 && obstacleGrid[i][j-1] == 0 {
				obstacleGrid[i][j] = 0
				continue
			}
			if j == 0 && i > 0 && obstacleGrid[i-1][j] == 0 {
				obstacleGrid[i][j] = 0
				continue
			}
			if obstacleGrid[i][j] == pointStone {
				obstacleGrid[i][j] = 0
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}
	return obstacleGrid[m-1][n-1]
}
