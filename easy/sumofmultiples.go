package main

// 2652. 倍数求和
// https://leetcode.cn/problems/sum-multiples/?envType=daily-question&envId=2023-10-17
// 给你一个正整数 n ，请你计算在 [1，n] 范围内能被 3、5、7 整除的所有整数之和，返回一个整数，用于表示给定范围内所有满足约束条件的数字之和。

// NOTICE: 等差数列求和，减掉重复的数值和，但是 n = 101 测试未通过，但是思路我感觉是没有错的
func sumOfMultiples(n int) int {
	max3, max5, max7 := n/3, n/5, n/7
	max15, max21, max35 := n/15, n/21, n/35
	return ((3+max3*3)*max3 + (5+max5*5)*max5 + (7+max7*7)*max7 - (15+max15*15)*max15 - (21+max21*21)*max21 - (35+max35*35)*max35) / 2
}

// 常规思路做法
func sumOfMultiples2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}
