package main

import "math/big"

/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？*/

// uniquePaths 62. 不同路径
func uniquePaths(m int, n int) int {
	var res int
	res = 0
	if n < 2 && m >= 2 {
		checkRecursionPath(2, 1, &res, m, n)
	}
	if m < 2 && n >= 2 {
		checkRecursionPath(1, 2, &res, m, n)
	}
	if m >= 2 && n >= 2 {
		checkRecursionPath(2, 1, &res, m, n)
		checkRecursionPath(1, 2, &res, m, n)
	}

	return res
}

func checkRecursionPath(startRow, startCol int, res *int, maxRow, maxCol int) {
	if startRow < maxRow && startCol < maxCol {
		*res = *res + 2
		checkRecursionPath(startRow+1, startCol, res, maxRow, maxCol)
		checkRecursionPath(startRow, startCol+1, res, maxRow, maxCol)
	}
	if startRow < maxRow && startCol >= maxCol {
		*res++
		checkRecursionPath(startRow+1, startCol, res, maxRow, maxCol)
	}
	if startRow >= maxRow && startCol < maxCol {
		*res++
		checkRecursionPath(startRow, startCol+1, res, maxRow, maxCol)
	}

	return
}

// 排列组合问题 C(m-1)/(m+n-2) 超出 memory limit
// Cn(m) = !n/(!m*!(n-m))
func uniquePaths2(m int, n int) int {
	return factorial(m+n-2) / (factorial(n-1) * factorial(m-1))
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// 官方答案
func uniquePaths3(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
