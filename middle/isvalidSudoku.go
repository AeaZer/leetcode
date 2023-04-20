package main

import "bytes"

func isValidSudoku(board [][]byte) bool {
	length := len(board)
	comb := make([]byte, 0, length)
	res := true
	var dfs func(row int)
	dfs = func(row int) {
		if (!res) || row == length {
			return
		}
		comb = append(comb, append([]byte(nil), board[row][:]...)...)
		check(comb, &res)
		comb = []byte{}
		for i := 0; i < length; i++ {
			comb = append(comb, board[i][row])
		}
		check(comb, &res)
		comb = []byte{}
		dfs(row + 1)
	}
	dfs(0)

	if !res {
		return false
	}

	comb = make([]byte, 0, 9)
	dfs = func(row int) {
		if (!res) || row == length/3 {
			return
		}

		for j := 0; j < 3; j++ {
			if !res {
				break
			}
			for i := row * 3; i < (row+1)*3; i++ {
				comb = append(comb, board[i][j*3:(j+1)*3]...)
			}
			check(comb, &res)
			comb = []byte{}
		}

		dfs(row + 1)
	}
	dfs(0)

	return res
}

func check(comb []byte, res *bool) {
	if !(*res) {
		return
	}
	for index, element := range comb {
		if !(*res) {
			break
		}
		if element == '.' {
			continue
		}
		if bytes.LastIndexByte(comb, element) != index {
			*res = false
		}
	}
}
