package string_array

/*给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，
后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。*/

/*输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。*/

// https://leetcode.cn/problems/merge-sorted-array/

// 向左移动，空出槽点
func rightMoveArr(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := len(arr) - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	var (
		remain = n // num2 未插入到 num1 的数量
		jf     int // 优化 num1 的遍历，因为 num2 其实是有序的
	)
	for i := 0; i < n; i++ {
		for j := jf; j < len(nums1); j++ {
			if nums2[i] < nums1[j] {
				rightMoveArr(nums1[j:])
				nums1[j] = nums2[i]
				remain--
				jf = j
				break
			}
		}
	}
	// 未插入的数据全部 copy 到 num1 的结尾
	if remain > 0 {
		copy(nums1[len(nums1)-remain:], nums2[n-remain:])
	}
}
