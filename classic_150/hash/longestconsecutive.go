package hash

/*给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。*/

// https://leetcode.cn/problems/longest-consecutive-sequence/

// 思路：先进行排序，然后对排序后的数组挨个遍历，每次遍历的长度如果大于最大长度那就替换当前长度为最大长度
func longestConsecutive(nums []int) int {
	numsL := len(nums)
	if numsL == 0 {
		return 0
	}
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	var maxLength int
	for i := 0; i < numsL; i++ {
		n, currentLength := nums[i], 1
		// 一直循环，直到不连续为止
		for {
			n++
			if !m[n] {
				if currentLength > maxLength {
					maxLength = currentLength
				}
				break
			}
			currentLength++
		}
	}

	return maxLength
}
