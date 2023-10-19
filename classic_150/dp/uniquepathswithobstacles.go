package dp

// 63. 不同路径 II
// https://leetcode.cn/problems/unique-paths-ii/description/?envType=study-plan-v2&envId=top-interview-150

/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[rows-1][cols-1] == 1 {
		return 0
	}
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = 1
	for i := 1; i < rows; i++ {
		// 是障碍物
		if obstacleGrid[i][0] == 1 {
			continue
		}
		dp[i][0] += dp[i-1][0]
	}
	for j := 1; j < cols; j++ {
		// 是障碍物
		if obstacleGrid[0][j] == 1 {
			continue
		}
		dp[0][j] += dp[0][j-1]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			// 障碍物
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[rows-1][cols-1]
}
