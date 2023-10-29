package _0231029

/*给你两个由正整数和 0 组成的数组 nums1 和 nums2 。

你必须将两个数组中的 所有 0 替换为 严格 正整数，并且满足两个数组中所有元素的和 相等 。

返回 最小 相等和 ，如果无法使两数组相等，则返回 -1 。*/

/*输入：nums1 = [3,2,0,1,0], nums2 = [6,5,0]
输出：12
解释：可以按下述方式替换数组中的 0 ：
- 用 2 和 4 替换 nums1 中的两个 0 。得到 nums1 = [3,2,2,1,4] 。
- 用 1 替换 nums2 中的一个 0 。得到 nums2 = [6,5,1] 。
两个数组的元素和相等，都等于 12 。可以证明这是可以获得的最小相等和。*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumAndZeroCount(nums []int) (sum int, zeroCount int) {
	for i := range nums {
		sum += nums[i]
		if nums[i] == 0 {
			zeroCount++
		}
	}
	return
}

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, c1 := sumAndZeroCount(nums1)
	sum2, c2 := sumAndZeroCount(nums2)

	if sum1 == sum2 {
		if (c1 == 0 && c2 != 0) || (c2 == 0 && c1 != 0) {
			return -1
		}
		return int64(sum1 + max(c1, c2))
	}
	if sum1 > sum2 {
		if c2 == 0 || (sum1-sum2 < c2 && c1 == 0) {
			return -1
		}
		if c1 < c2 && (c2-(sum1-sum2) > c1) {
			return int64(sum2 + c1)
		}
		return int64(sum1 + c1)
	} else {
		if c1 == 0 || (sum2-sum1 < c1 && c2 == 0) {
			return -1
		}
		if c1 > c2 && (c1-(sum2-sum1) > c2) {
			return int64(sum1 + c1)
		}
		return int64(sum2 + c2)
	}
}
