package _0240303

func resultArray(nums []int) []int {
	tmp1, tmp2 := make([]int, 1), make([]int, 1)
	tmp1[0], tmp2[0] = nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		if tmp1[len(tmp1)-1] > tmp2[len(tmp2)-1] {
			tmp1 = append(tmp1, nums[i])
			continue
		}
		tmp2 = append(tmp2, nums[i])
	}
	ans := make([]int, 0, len(nums))
	ans = append(tmp1, tmp2...)
	return ans
}

type keyPairs struct {
	col, row int
}

func countSubmatrices(grid [][]int, k int) int {
	var ans int
	n, m := len(grid), len(grid[0])
	dp := make(map[keyPairs]int, n*m)
	var firstRowSum int
	for i := 0; i < m; i++ {
		firstRowSum += grid[0][i]
		if firstRowSum <= k {
			ans++
		}
		dp[keyPairs{col: 0, row: i}] = firstRowSum
	}
	for i := 1; i < n; i++ {
		var currentRowSum int
		for j := 0; j < m; j++ {
			currentRowSum += grid[i][j]
			tmp := dp[keyPairs{col: i - 1, row: j}] + currentRowSum
			if tmp <= k {
				ans++
			}
			dp[keyPairs{col: i, row: j}] = tmp
		}
	}
	return ans
}
