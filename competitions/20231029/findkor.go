package _0231029

import "math"

/*给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

nums 中的 K-or 是一个满足以下条件的非负整数：

只有在 nums 中，至少存在 k 个元素的第 i 位值为 1 ，那么 K-or 中的第 i 位的值才是 1 。
返回 nums 的 K-or 值。

注意 ：对于整数 x ，如果 (2i AND x) == 2i ，则 x 中的第 i 位值为 1 ，其中 AND 为按位与运算符。*/

/*输入：nums = [7,12,9,8,9,15], k = 4
输出：9
解释：nums[0]、nums[2]、nums[4] 和 nums[5] 的第 0 位的值为 1 。
nums[0] 和 nums[5] 的第 1 位的值为 1 。
nums[0]、nums[1] 和 nums[5] 的第 2 位的值为 1 。
nums[1]、nums[2]、nums[3]、nums[4] 和 nums[5] 的第 3 位的值为 1 。
只有第 0 位和第 3 位满足数组中至少存在 k 个元素在对应位上的值为 1 。因此，答案为 2^0 + 2^3 = 9 。*/

/*1 <= nums.length <= 50
0 <= nums[i] < 2^31
1 <= k <= nums.length*/

func findKOr(nums []int, k int) int {
	bits := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if 1<<i > num {
				break
			}
			if num&(1<<i) != 0 {
				bits[i]++
			}
		}
	}
	var ans int
	for i := 0; i < 32; i++ {
		if bits[i] >= k {
			ans += int(math.Pow(2, float64(i)))
		}
	}
	return ans
}
