package math

// 66. 加一
// https://leetcode.cn/problems/plus-one/?envType=study-plan-v2&envId=top-interview-150

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。

func plusOne(digits []int) []int {
	rest := 1
	for i := len(digits) - 1; i > -1; i-- {
		sum := digits[i] + rest
		digits[i] = sum % 10
		rest = sum / 10
	}
	if rest != 0 {
		digits = append([]int{rest}, digits...)
	}
	return digits
}
