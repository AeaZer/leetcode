package matrix

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	tmp := make([][]int, n+2)
	for i := range tmp {
		tmp[i] = make([]int, m+2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp[i][j] = board[i+1][j+1]
		}
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {

		}
	}
}
