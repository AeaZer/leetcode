package graph

// 130. 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/

/*给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。*/

const (
	elementSingedX       byte = 'X'
	elementSingedO       byte = 'O'
	elementSingedVisited byte = 'V'
)

func solve(board [][]byte) {
	iMax, jMax := len(board), len(board[0])
	var dfs func(rowIndex, columnIndex int)
	dfs = func(rowIndex, columnIndex int) {
		if rowIndex >= iMax || rowIndex < 0 || columnIndex >= jMax || columnIndex < 0 ||
			board[rowIndex][columnIndex] != elementSingedO {
			return
		}
		board[rowIndex][columnIndex] = elementSingedVisited

		dfs(rowIndex+1, columnIndex) // 右边的节点
		dfs(rowIndex-1, columnIndex) // 左边的节点
		dfs(rowIndex, columnIndex+1) // 下边的节点
		dfs(rowIndex, columnIndex-1) // 上面的节点
	}

	for i := 0; i < iMax; i++ {
		dfs(i, 0)      // left
		dfs(i, jMax-1) // right
	}
	for j := 0; j < jMax; j++ {
		dfs(0, j)      // top
		dfs(iMax-1, j) // bottom
	}

	for i := 0; i < iMax; i++ {
		for j := 0; j < jMax; j++ {
			switch board[i][j] {
			case elementSingedO:
				board[i][j] = elementSingedX
			case elementSingedVisited:
				board[i][j] = elementSingedO
			}
		}
	}

}
