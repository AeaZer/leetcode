package matrix

/*请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一*/

/*注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。*/

// https://leetcode.cn/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var (
		rows     [9][9]uint8    // 行
		columns  [9][9]uint8    // 列
		subBoxes [3][3][9]uint8 // 子数独
	)
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subBoxes[i/3][j/3][index]++
			// 分别是每一行、每一列、每个子 3*3 数独中的校验
			if rows[i][index] > 1 || columns[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
