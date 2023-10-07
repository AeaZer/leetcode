package recall

const (
	stateFail = iota
	stateNext
)

func exist(board [][]byte, word string) bool {
	wordBytes := []byte(word)
	var (
		bl     = len(wordBytes)
		colL   = len(board)
		rowL   = len(board[0])
		found  bool
		search func(int, int, int, int)
	)
	if bl > colL*rowL {
		return false
	}
	visited := make([][]bool, colL)
	for i := range visited {
		visited[i] = make([]bool, rowL)
	}
	search = func(i, j, index, state int) {
		if found {
			return
		}
		if board[i][j] == wordBytes[index] {
			if index == bl-1 {
				found = true
				return
			}
			// 向左查找
			if j-1 >= 0 && !visited[i][j-1] {
				visited[i][j-1] = true
				search(i, j-1, index+1, stateNext)
				visited[i][j-1] = false
			}

			// 向右查找
			if j+1 < rowL && !visited[i][j+1] {
				visited[i][j+1] = true
				search(i, j+1, index+1, stateNext)
				visited[i][j+1] = false
			}

			// 向下查找
			if i+1 < colL && !visited[i+1][j] {
				visited[i+1][j] = true
				search(i+1, j, index+1, stateNext)
				visited[i+1][j] = false
			}

			// 向上查找
			if i-1 >= 0 && !visited[i-1][j] {
				visited[i-1][j] = true
				search(i-1, j, index+1, stateNext)
				visited[i-1][j] = false
			}
		}
		if state == stateNext {
			return
		}
		// 查找第一个
		if i+1 < colL {
			visited[i+1][j] = true
			search(i+1, j, 0, stateFail)
			visited[i+1][j] = false
		}
		if j+1 < rowL {
			visited[i][j+1] = true
			search(i, j+1, 0, stateFail)
			visited[i][j+1] = false
		}
	}
	search(0, 0, 0, stateFail)
	return found
}
