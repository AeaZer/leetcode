package binarysearch

// 74. 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-interview-150

/*给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非递减顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。*/

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
