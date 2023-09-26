package recall

func exist(board [][]byte, word string) bool {
	wordBytes := []byte(word)
	var (
		bl     = len(wordBytes)
		rowL   = len(board)
		colL   = len(board[0])
		found  bool
		search func(int, int, int)
	)
	search = func(rowIndex, colIndex, index int) {
		if found || rowIndex >= rowL || colIndex >= colL {
			return
		}
		if board[rowIndex][colIndex] == wordBytes[index] {
			if index == bl-1 {
				found = true
				return
			}
			search(rowIndex, colIndex-1, index+1)
			search(rowIndex, colIndex+1, index+1)
			search(rowIndex+1, colIndex, index+1)
		}
		search(rowIndex, colIndex+1, index)
		search(rowIndex+1, colIndex, index)
	}
	return found
}
