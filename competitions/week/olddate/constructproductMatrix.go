package olddate

func constructProductMatrix(grid [][]int) [][]int {
	multiTotal := 1
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			multiTotal *= grid[i][j]
		}
	}
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = (multiTotal / grid[i][j]) % 12345
		}
	}
	return ans
}
