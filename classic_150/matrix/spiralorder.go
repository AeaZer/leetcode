package matrix

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

// https://leetcode.cn/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	var (
		n           = len(matrix)
		m           = len(matrix[0])
		nMultiM     = m * n
		res         = make([]int, 0, nMultiM)
		count       int
		i, j        int
		spiralTimes int
	)
	if n == 1 {
		return matrix[0]
	}
	if m == 1 {
		for z := 0; z < n; z++ {
			res = append(res, matrix[z][0])
		}
		return res
	}

	// 假设正在遍历 3*4
	for count < nMultiM-1 {
		// a[0][0], a[1][1]
		i, j = spiralTimes, spiralTimes
		for {
			if j > m-2-spiralTimes {
				break
			}
			res = append(res, matrix[i][j])
			j++
			count++
		}
		// a[0][3], a[1][2]
		i, j = spiralTimes, m-1-spiralTimes
		for {
			if i > n-2-spiralTimes {
				break
			}
			res = append(res, matrix[i][j])
			i++
			count++
		}
		// a[2][3], a[1][2]
		i, j = n-1-spiralTimes, m-1-spiralTimes
		for {
			if j <= spiralTimes {
				break
			}
			res = append(res, matrix[i][j])
			j--
			count++
		}
		// a[2][3], a[1][2]
		i, j = n-1-spiralTimes, spiralTimes
		for {
			if i <= spiralTimes {
				break
			}
			res = append(res, matrix[i][j])
			i--
			count++
		}
		spiralTimes++
	}
	if count == nMultiM-1 {
		res = append(res, matrix[n/2][m/2])
	}
	return res
}
