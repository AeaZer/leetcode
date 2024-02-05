package dp

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	x, y, z, w := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		x = y
		y = z
		z = w
		w = x + y + z
	}
	return w
}
