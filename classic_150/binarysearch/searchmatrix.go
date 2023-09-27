package binarysearch

// [1,3,5,7],[10,11,16,20],[23,30,34,60]

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}

	i := len(matrix) - 1
	for rowIndex := 0; rowIndex < len(matrix)-1; rowIndex++ {
		if matrix[rowIndex][0] <= target && matrix[rowIndex+1][0] > target {
			i = rowIndex
			break
		}
	}
	rowLeft, rowRight := 0, len(matrix[0])-1
	for {
		mid := (rowLeft + rowRight) / 2
		if matrix[i][mid] == target {
			return true
		}
		if rowLeft == mid {
			if matrix[i][rowRight] == target {
				return true
			}
			return false
		}
		if matrix[i][mid] > target {
			rowRight = mid
			continue
		}
		rowLeft = mid
	}
}
