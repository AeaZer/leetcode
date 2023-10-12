package main

import (
	"math"
	"strconv"
)

// 2562. 找出数组的串联值
// https://leetcode.cn/problems/find-the-array-concatenation-value/description/

/*给你一个下标从 0 开始的整数数组 nums 。

现定义两个数字的 串联 是由这两个数值串联起来形成的新数字。

例如，15 和 49 的串联是 1549 。
nums 的 串联值 最初等于 0 。执行下述操作直到 nums 变为空：

如果 nums 中存在不止一个数字，分别选中 nums 中的第一个元素和最后一个元素，将二者串联得到的值加到 nums 的 串联值 上，然后从 nums 中删除第一个和最后一个元素。
如果仅存在一个元素，则将该元素的值加到 nums 的串联值上，然后删除这个元素。
返回执行完所有操作后 nums 的串联值。*/

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	nl := len(nums)
	left, right := nl/2-1, nl/2
	if nl&1 != 0 {
		res = int64(nums[right])
		right++
	}
	for left >= 0 {
		res += cascade(nums[left], nums[right])
		left--
		right++
	}
	return res
}

func cascade(n1, n2 int) int64 {
	return int64(n1*int(math.Pow10(len(strconv.Itoa(n2)))) + n2)
}
