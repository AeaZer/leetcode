package recall

// 46. 全排列
// https://leetcode.cn/problems/permutations/

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 额外数组辅助
func permute(nums []int) [][]int {
	var (
		nl      = len(nums)
		res     = make([][]int, 0)
		visited = make([]bool, nl) // 额外数组
	)
	var backtrack func([]int)
	backtrack = func(path []int) {
		if len(path) == nl {
			comb := make([]int, nl)
			copy(comb, path)
			res = append(res, comb)
			return
		}
		for i := 0; i < nl; i++ {
			if !visited[i] {
				visited[i] = true
				path = append(path, nums[i])
				backtrack(path)
				visited[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack([]int{})
	return res
}

// 递归交换
func permute2(nums []int) [][]int {
	var (
		res       = make([][]int, 0)
		backtrack func(int)
	)

	backtrack = func(index int) {
		if index == len(nums) {
			// 将当前排列添加到结果集合
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := index; i < len(nums); i++ {
			// 交换元素位置
			nums[index], nums[i] = nums[i], nums[index]
			// 进入下一层递归
			backtrack(index + 1)
			// 恢复原始顺序
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	backtrack(0)
	return res
}
