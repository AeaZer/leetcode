package bitcup

// 191. 位1的个数
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
// https://leetcode.cn/problems/number-of-1-bits/description/?envType=study-plan-v2&envId=top-interview-150

func hammingWeight(num uint32) int {
	var c int
	for i, n := 0, uint32(1); i < 32; i++ {
		if num&n != 0 {
			c++
		}
		n <<= 1
	}
	return c
}
