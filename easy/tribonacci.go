package main

/*泰波那契序列 Tn 定义如下：
T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
给你整数 n，请返回第 n 个泰波那契数 Tn 的值*/

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	x, y, z, temp := 0, 1, 1, 0
	for i := 0; i < n-2; i++ {
		temp = z
		z = x + y + z
		x = y
		y = temp
	}

	return z
}
