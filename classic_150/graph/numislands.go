package graph

// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/

/*给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。*/

/*m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'*/

const (
	spaceWater   byte = '0'
	spaceLand    byte = '1'
	spaceVisited byte = '2'
)

func numIslands(grid [][]byte) int {
	var res int
	iMax, jMax := len(grid), len(grid[0])
	var dfs func(rowIndex, columnIndex int)
	dfs = func(rowIndex, columnIndex int) {
		// 下标不在给定的区域内
		if !inArea(iMax, jMax, rowIndex, columnIndex) {
			return
		}
		// 如果此节点是"水"或者"已被访问过"那就跳出递归
		if grid[rowIndex][columnIndex] == spaceWater || grid[rowIndex][columnIndex] == spaceVisited {
			return
		}
		grid[rowIndex][columnIndex] = spaceVisited // grid[rowIndex][columnIndex] == spaceLand
		dfs(rowIndex+1, columnIndex)               // 右边的节点
		dfs(rowIndex-1, columnIndex)               // 左边的节点
		dfs(rowIndex, columnIndex+1)               // 下边的节点
		dfs(rowIndex, columnIndex-1)               // 上面的节点
	}
	for i := 0; i < iMax; i++ {
		for j := 0; j < jMax; j++ {
			if grid[i][j] == spaceLand {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

func inArea(iMax, jMax, rowIndex, columnIndex int) bool {
	return rowIndex < iMax && rowIndex >= 0 && columnIndex < jMax && columnIndex >= 0
}
